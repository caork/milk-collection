package LinkedRoute

import (
	"math"
	"math/rand"
	"milk-collection/src/math32"
	"milk-collection/src/parser"
	"milk-collection/src/routeStructure"
)

func remove(s []int, i int) []int { // order is not matters
	var ind int
	for index, value := range s {
		if value == i {
			ind = index
			break
		}
	}
	s[ind] = s[len(s)-1]
	return s[:len(s)-1]
}

func totalDistance(route *routeStructure.List, records *parser.LocationRecords) float32 {
	if route.Len() <= 1 {
		return math.MaxFloat32
	}
	first := route.Front().Value
	from := route.Front().Value
	var distance float32 = 0.0
	for p := route.Front().Next(); p != nil; p = p.Next() {
		distance += records.Distance(from, p.Value)
		from = p.Value
	}
	distance += records.Distance(from, first)
	return distance
}

func reArrange(currentPlace *routeStructure.List, placeNotIn []int, records *parser.LocationRecords) (int, *routeStructure.Element) {

	var place int
	var insertAfter *routeStructure.Element
	var currentCost float32
	var minimumCost float32 = math.MaxFloat32
	for _, p := range placeNotIn {
		for i := currentPlace.Front(); i != nil; i = i.Next() {
			if currentPlace.Len() == 1 {
				currentCost = records.Distance(i.Value, p)
			} else {
				j := i.Next() // next one
				if j == nil { // if i in the tail
					j = currentPlace.Front()
				}
				currentCost = cost(i.Value, j.Value, p, records)
			}
			if currentCost < minimumCost {
				minimumCost = currentCost
				place = p
				insertAfter = i
			}
		}
	}
	return place, insertAfter
}

func cost(left int, right int, insertValue int, records *parser.LocationRecords) float32 {
	preCost := records.Distance(left, right)
	currentCost := records.Distance(left, insertValue) + records.Distance(insertValue, right)
	return math32.Abs(currentCost - preCost)
}

func CheapestInsert(records *parser.LocationRecords, startPlace int) (*routeStructure.List, float32) {
	if startPlace == -1 {
		startPlace = rand.Intn(len(records.Index))
	}

	placeNotIn := make([]int, len(records.Index))
	copy(placeNotIn, records.Index)

	currentPlace := routeStructure.New()
	currentPlace.PushBack(startPlace)
	placeNotIn = remove(placeNotIn, startPlace)

	for len(placeNotIn) >= 1 {
		chosenPlace, index := reArrange(currentPlace, placeNotIn, records)
		currentPlace.InsertAfter(chosenPlace, index)
		placeNotIn = remove(placeNotIn, chosenPlace)
	}
	return currentPlace, totalDistance(currentPlace, records)
}
