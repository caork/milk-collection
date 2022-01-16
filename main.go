package main

import (
	"milk-collection/LinkedRoute"
	"milk-collection/parser"
	"os"
	"runtime/pprof"
)

func main() {
	f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	df := parser.TableLoader("finland1000.csv")
	// Summarize(&df, NearestNeighbor, "NearestNeighbor")
	// Summarize(&df, NearestInsert, "NearestInsert")
	// parser.Summarize(&df, sliceRoute.FarthestInsert, "FarthestInsert") // 43799.413274
	parser.SummarizeLinked(&df, LinkedRoute.CheapestInsert, "CheapestInsert")
}
