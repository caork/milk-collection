package main

import (
	"encoding/csv"
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

func main() {
	df := tableLoader("distance.csv")
	for _, s := range NearestNeighbor(&df) {
		println(s)
	}
}
