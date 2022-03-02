package main

import (
	"milk-collection/src/LinkedRoute"
	"milk-collection/src/parser"
	"os"
	"runtime/pprof"
)

func main() {
	fc, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer fc.Close()
	pprof.StartCPUProfile(fc)
	defer pprof.StopCPUProfile()

	fm, _ := os.OpenFile("mem.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer fm.Close()

	pprof.Lookup("heap").WriteTo(fm, 0)

	df := parser.TableLoader("data/finland1000.csv")
	// Summarize(&df, NearestNeighbor, "NearestNeighbor")
	// Summarize(&df, NearestInsert, "NearestInsert")
	// parser.Summarize(&df, sliceRoute.FarthestInsert, "FarthestInsert") // 43799.413274
	parser.SummarizeLinked(&df, LinkedRoute.CheapestInsert, "CheapestInsert") //41577.380091
}
