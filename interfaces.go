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
	jade.SetWorkerModule(worker)

	aggregator := digest.NewStomach()
	jade.SetAggregatorModule(aggregator)
	jade.CreateHTTPServer()
}
