package sliceRoute

import (
	"math"
	"math/rand"
	"milk-collection/src/parser"
)

func remove(s []int, i int) []int { // order is not matters
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeWithOrder(s []int, i int) []int { // order matters
	return append(s[:i], s[i+1:]...)
}

func totalDistance(s []int, records *parser.LocationRecords) float32 { // total distance
	if len(s) <= 1 {
		return math.MaxFloat32
	}
	first := s[0]
	from := s[0]
	var total float32 = 0.0
	for i := 1; i < len(s); i++ {
		total += records.Distance(from, s[i])
		from = s[i]
	}
	total += records.Distance(from, first) // become a loop
	return total
}

func NearestNeighbor(records *parser.LocationRecords) []int {
	var currentPlace = make([]int, 0)
	placesNotIn := make([]int, len(records.Index))
	copy(placesNotIn, records.Index)
	chosenPlace := rand.Intn(len(placesNotIn))

	nextPlace := func() int {
		nearestIndex := 0 // get the index of value
		nearestId := placesNotIn[0]
		nearest := records.Distance(chosenPlace, placesNotIn[0]) // use the first as default the nearest distance
		for i := 1; i < len(placesNotIn); i++ {
			distance := records.Distance(chosenPlace, placesNotIn[i])
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

func FarthestInsert(records *parser.LocationRecords) []int {
	var distance float32 = math.MaxFloat32
	route := make([]int, 0)
	for _, s := range records.Index {
		thisRoute := NearestOrFarthestInsert(records, s)
		thisWay := totalDistance(thisRoute, records)
		if thisWay < distance {
			distance = thisWay
			route = thisRoute
		}
	}
	return route
}

func NearestOrFarthestInsert(records *parser.LocationRecords, startPlace int) []int {
	var currentPlace = make([]int, 0)
	placesNotIn := make([]int, len(records.Index))
	copy(placesNotIn, records.Index)
	chosenPlace := startPlace
	currentPlace = append(currentPlace, chosenPlace)
	placesNotIn = remove(placesNotIn, chosenPlace)

	reArrangeRoute := func(newPlace int) ([]int, float32) {
		var shortestDistance float32 = math.MaxFloat32
		var shortestRoute []int //var shortestRoute = make([]int, len(records.index))
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
		return shortestRoute, shortestDistance

	}
	var shortestRoute []int
	var shortestDistance float32 = math.MaxFloat32
	var theChoosePlace int
	for len(placesNotIn) >= 1 {
		for i, p := range placesNotIn {
			currentRoute, currentDistance := reArrangeRoute(p)
			if currentDistance < shortestDistance {
				currentDistance = shortestDistance
				shortestRoute = currentRoute
				theChoosePlace = i
			}
		}
		currentPlace = shortestRoute
		placesNotIn = remove(placesNotIn, theChoosePlace)

	}
	return currentPlace

}
