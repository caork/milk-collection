package main

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

type locationRecords struct {
	records      [][]string
	index        []int
	distance     func(int, int) float64
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

	location := locationRecords{
		records: records,
		index:   NewSlice(1, len(records), 1),
		distance: func(from int, to int) float64 {

			getFromIndex := func(index int) [2]int64 {
				var thisLocation [2]int64
				info := records[index]
				thisLocation[0], _ = strconv.ParseInt(info[1], 0, 64)
				thisLocation[1], _ = strconv.ParseInt(info[2], 0, 64)
				return thisLocation
			}
			fromLocation := getFromIndex(from)
			toLocation := getFromIndex(to)
			horizontal := fromLocation[0]-toLocation[0]
			vertical := fromLocation[1]-toLocation[1]
			return math.Sqrt(math.Pow(float64(horizontal),2)+math.Pow(float64(vertical),2))
		},
	}
	return location
}

func NewSlice(start, end, step int) []int {
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
	println(df.distance(3,4))
}


