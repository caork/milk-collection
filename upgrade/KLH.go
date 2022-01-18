package upgrade

import (
	"github.com/RoaringBitmap/roaring"
	"math"
	"math/rand"
	"milk-collection/parser"
	"milk-collection/routeStructure"
	"time"
)

func exchangeCost(records *parser.LocationRecords, leftPoint int, rightPoint int) float32 {
	originDistance := records.Distance(leftPoint-1, leftPoint) +
		records.Distance(leftPoint, leftPoint+1) +
		records.Distance(rightPoint-1, rightPoint) +
		records.Distance(rightPoint, rightPoint+1)
	newDistance := records.Distance(leftPoint-1, rightPoint) +
		records.Distance(rightPoint, leftPoint+1) +
		records.Distance(rightPoint-1, leftPoint) +
		records.Distance(leftPoint, rightPoint+1)
	return newDistance - originDistance
}

func newPair(bitmap *roaring.Bitmap, routeRange uint32, records *parser.LocationRecords) (uint32, uint32) {
	rand.Seed(time.Now().UnixNano())
	var left, right uint32
	var times int8 = 0
	var pair uint32
	for times <= math.MaxInt8-1 {
		times++
		left = rand.Uint32() % routeRange
		right = rand.Uint32() % routeRange
		pair = left * 103 & right
		if !bitmap.Contains(pair) { // avoid most duplicate distance calculate
			if exchangeCost(records, int(left), int(right)) < 0 {
				bitmap.Add(pair)
				break
			}
		}
	}
	return left, right
}

func annealing(route *routeStructure.List, records *parser.LocationRecords) {
	sliceRoute := route.ToSlice()
	alreadyChosen := roaring.New()
	exchange := func(left int, right int) {
		tmp := sliceRoute[left]
		sliceRoute[left] = sliceRoute[right]
		sliceRoute[right] = sliceRoute[tmp]
	}

	for true {
		left, right := newPair(alreadyChosen, uint32(route.Len()), records)
		exchange(int(left), int(right))
	}

}
