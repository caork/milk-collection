package main

import (
	"milk-collection/LinkedRoute"
	"milk-collection/parser"
)

func main() {
	df := parser.TableLoader("finland1000.csv")
	// Summarize(&df, NearestNeighbor, "NearestNeighbor")
	// Summarize(&df, NearestInsert, "NearestInsert")
	// parser.Summarize(&df, sliceRoute.FarthestInsert, "FarthestInsert") // 43799.413274
	parser.SummarizeLinked(&df, LinkedRoute.CheapestInsert, "CheapestInsert")
}
