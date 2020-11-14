package main

import (
	"aces/plankton/digest"
	"aces/plankton/eat"
	"uta.edu/aces/jadesdk"
)

func main() {
	jade := jadesdk.NewJadeSDK()
	jade.Verbose(true)

	worker := eat.NewMouth()
	jade.SetDefaultWorkerModule(worker)

	aggregator := digest.NewStomach()
	jade.SetDefaultAggregatorModule(aggregator)
	jade.CreateHTTPServer()
}
