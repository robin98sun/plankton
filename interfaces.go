package main

import (
	"aces/plankton/digest"
	"aces/plankton/eat"
	"uta.edu/aces/jadesdk"
	"runtime/debug"
)

func main() {
	jade := jadesdk.NewJadeSDK()
	jade.Verbose(false)
	// disable GC at runtime
	debug.SetGCPercent(-1)

	worker := eat.NewMouth()
	jade.SetDefaultWorkerModule(worker)

	aggregator := digest.NewStomach()
	jade.SetDefaultAggregatorModule(aggregator)
	jade.CreateHTTPServer()
}
