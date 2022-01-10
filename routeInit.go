package main

import (
	"math"
	"math/rand"
)

type routeAlgorithm func(records *locationRecords) []int // interface of route algorithms

func remove(s []int, i int) []int { // order is not matters
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeWithOrder(s []int, i int) []int { // order matters
	return append(s[:i], s[i+1:]...)
}

func totalDistance(s []int, records *locationRecords) float64 { // total distance
	if len(s) <= 1 {
		return math.MaxFloat64
	}
	first := s[0]
	from := s[0]
	var total float64 = 0.0
	for i := 1; i < len(s); i++ {
		total += records.distance(from, s[i])
		from = s[i]
	}
	total += records.distance(from, first) // become a loop
	return total
}

func NearestNeighbor(records *locationRecords) []int {
	var currentPlace = make([]int, 0)
	placesNotIn := make([]int, len(records.index))
	copy(placesNotIn, records.index)
	chosenPlace := rand.Intn(len(placesNotIn))

	nextPlace := func() int {
		nearestIndex := 0 // get the index of value
		nearestId := placesNotIn[0]
		nearest := records.distance(chosenPlace, placesNotIn[0]) // use the first as default the nearest distance
		for i := 1; i < len(placesNotIn); i++ {
			distance := records.distance(chosenPlace, placesNotIn[i])
			if distance < nearest {
				nearestId = placesNotIn[i]
				nearest = distance
				nearestIndex = i
			}
		}
		placesNotIn = remove(placesNotIn, nearestIndex)
		return nearestId
	}

	for len(placesNotIn) > 1 { // append the nearest one
		next := nextPlace()
		currentPlace = append(currentPlace, next)
		chosenPlace = next
	}
	return append(currentPlace, placesNotIn[0]) // add the last one
}

func NearestInsert(records *locationRecords) []int {
	var currentPlace = make([]int, 0)
	placesNotIn := make([]int, len(records.index))
	copy(placesNotIn, records.index)
	chosenPlace := rand.Intn(len(placesNotIn))
	currentPlace = append(currentPlace, chosenPlace)
	placesNotIn = remove(placesNotIn, chosenPlace)
	isFirstChoose := true
	choose := func() int {
		var nearest float64
		if isFirstChoose {
			nearest = records.distance(chosenPlace, placesNotIn[0])
			isFirstChoose = false
		} else {
			nearest = math.MaxFloat64
		}
		targetIndex := 0
		for _, place := range currentPlace {
			for j, target := range placesNotIn {
				distance := records.distance(place, target)
				if distance < nearest {
					nearest = distance
					targetIndex = j
				}
			}
		}
		returnIndex := placesNotIn[targetIndex]
		placesNotIn = remove(placesNotIn, targetIndex)
		return returnIndex
	}

	reArrangeRoute := func(newPlace int) {
		shortestDistance := math.MaxFloat64
		var shortestRoute = make([]int, len(records.index))
		var fairlyShortRoute []int
		var currentPlaceCopy = make([]int, len(currentPlace))
		copy(currentPlaceCopy, currentPlace)
		for i := range currentPlaceCopy {
			if i == 0 {
				fairlyShortRoute = append([]int{newPlace}, currentPlaceCopy...)
			} else if i == len(currentPlaceCopy)-1 {
				fairlyShortRoute = append(currentPlaceCopy, newPlace)
			} else {
				fairlyShortRoute = append(currentPlaceCopy[:i+1], currentPlaceCopy[i:]...)
				fairlyShortRoute[i] = newPlace
			}
			distance := totalDistance(fairlyShortRoute, records)
			if shortestDistance > distance {
				shortestDistance = distance
				shortestRoute = fairlyShortRoute
			}
		}
		currentPlace = shortestRoute

	}

	for len(placesNotIn) > 1 {
		reArrangeRoute(choose())
	}
	reArrangeRoute(placesNotIn[0])
	return currentPlace

}
