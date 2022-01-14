package parser

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"milk-collection/routeStructure"
	"os"
	"strconv"
	"strings"
)

type RouteAlgorithm func(records *LocationRecords) []int // interface of route algorithms
type RouteAlgorithmLinked func(records *LocationRecords, chosePlace int) (*routeStructure.List, float64)

type LocationRecords struct {
	records       [][]string
	Index         []int
	storedRecords [][]float64
	Distance      func(int, int) float64
}

func LinkedToString(route *routeStructure.List, stringRoute *strings.Builder) *strings.Builder {
	stringRoute.Reset()
	isFirst := true
	for i := route.Front(); i != nil; i = i.Next() {
		if isFirst {
			stringRoute.WriteString(strconv.Itoa(i.Value))
			isFirst = false
		} else {
			stringRoute.WriteString("->")
			stringRoute.WriteString(strconv.Itoa(i.Value))
		}
	}
	return stringRoute

}

func TableLoader(filePath string) LocationRecords { // csv loader
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

	location := LocationRecords{
		records:       records,
		Index:         newSlice(0, len(records)-1, 1),
		storedRecords: storedRecords,
		Distance: func(from int, to int) float64 {

			getFromIndex := func(index int) [2]float64 {
				var thisLocation [2]float64
				info := records[index]
				thisLocation[0], _ = strconv.ParseFloat(info[1], 64)
				thisLocation[1], _ = strconv.ParseFloat(info[2], 64)
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

func TotalDistance(route []int, records *LocationRecords) float64 {
	distance := 0.0
	origin := -1
	from := -1
	for i, to := range route {
		if i == 0 {
			origin = to
			from = to
		} else if i < len(route)-1 {
			distance += records.Distance(from, to)
			from = to
		} else {
			distance += records.Distance(to, origin)
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

func Summarize(local *LocationRecords, algorithm RouteAlgorithm, name string) {
	route := algorithm(local)
	fmt.Println("Route Algorithm: ", name)
	RoutePrint(route)
	fmt.Println("")
	fmt.Printf("Toal distance is %f \n", TotalDistance(route, local))
}

func SummarizeLinked(local *LocationRecords, algorithm RouteAlgorithmLinked, name string) {
	route, distance := algorithm(local, 1)
	var routeline strings.Builder
	fmt.Println("Route Algorithm: ", name)
	routeline = *LinkedToString(route, &routeline)
	fmt.Println(routeline.String())
	fmt.Printf("Toal distance is %f \n", distance)
}
