package main

import (
	"aces/jadesdk"
	"aces/plankton/digest"
	"aces/plankton/eat"
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
