package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type locationRecords struct {
	records       [][]string
	index         []int
	storedRecords [][]float64
	distance      func(int, int) float64
}

func tableLoader(filePath string) locationRecords { // csv loader
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	storedRecords := newMatrix(len(records), len(records), -1)

	location := locationRecords{
		records:       records,
		index:         newSlice(0, len(records)-1, 1),
		storedRecords: storedRecords,
		distance: func(from int, to int) float64 {

			getFromIndex := func(index int) [2]int64 {
				var thisLocation [2]int64
				info := records[index]
				thisLocation[0], _ = strconv.ParseInt(info[1], 0, 64)
				thisLocation[1], _ = strconv.ParseInt(info[2], 0, 64)
				return thisLocation
			}

			if from == to {
				return math.MaxFloat64
			} else if storedRecords[from][to] != -1 {
				return storedRecords[from][to]
			} else {
				fromLocation := getFromIndex(from)
				toLocation := getFromIndex(to)
				horizontal := fromLocation[0] - toLocation[0]
				vertical := fromLocation[1] - toLocation[1]
				thisDistance := math.Sqrt(math.Pow(float64(horizontal), 2) + math.Pow(float64(vertical), 2))
				storedRecords[from][to] = thisDistance
				return thisDistance
			}
		},
	}
	return location
}

func newMatrix(row int, col int, fill float64) [][]float64 {
	m := make([][]float64, row)
	for i := range m {
		m[i] = make([]float64, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = fill
		}
	}
	return m
}

func newSlice(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

func TotalDistance(route []int, records *locationRecords) float64 {
	distance := 0.0
	origin := -1
	from := -1
	for i, to := range route {
		if i == 0 {
			origin = to
			from = to
		} else if i < len(route)-1 {
			distance += records.distance(from, to)
			from = to
		} else {
			distance += records.distance(to, origin)
		}
	}
	return distance
}

func RoutePrint(route []int) {
	origin := -1
	for i, s := range route {
		if i == 0 {
			origin = s
			fmt.Printf("%d -> ", s)
		} else if i < len(route)-1 {
			fmt.Printf("%d -> ", s)
		} else {
			fmt.Printf("%d -> %d", s, origin)
		}
	}
}

func Summarize(local *locationRecords, algorithm routeAlgorithm, name string) {
	route := algorithm(local)
	fmt.Println("Route Algorithm: ", name)
	RoutePrint(route)
	fmt.Println("")
	fmt.Printf("Toal distance is %f \n", TotalDistance(route, local))
}

func main() {
	df := tableLoader("distance.csv")
	Summarize(&df, NearestNeighbor, "NearestNeighbor")
	Summarize(&df, NearestInsert, "NearestInsert")
	Summarize(&df, FarthestInsert, "FarthestInsert")
}
