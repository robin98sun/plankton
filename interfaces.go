package main

import (
	"aces/video-comb-aggregator/kernel"
	"aces/video-comb-aggregator/utils"
	// RESTful Server
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"

	// others
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// Read command line flags
	port := flag.Int("port", 8080, "port to listen")
	addr := flag.String("addr", "0.0.0.0", "ip address to listen")
	flag.Parse()
	// Read Env configurations
	conf := utils.ReadConfFromEnv()
	if conf.SelfNode.Addr != "" {
		addr = &conf.SelfNode.Addr
	}
	if conf.SelfNode.Port != 0 {
		port = &conf.SelfNode.Port
	}
	confstr, _ := json.Marshal(conf)
	log.Println("Conf:", string(confstr))
	// Run the detection task
	core := kernel.NewCore(conf)

	// API responders
	status := func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(core.Status())
	}

	receive := func(w rest.ResponseWriter, r *rest.Request) {
		body, err := utils.DecodeRequestBody(r, nil)
		if err != nil {
			msg := "error when receiving result: " + err.Error()
			log.Println("[reducer]", msg)
			w.WriteJson(&utils.Response{
				Status: "ERROR",
				Error:  msg,
			})
		} else {
			core.ReceiveResults(body)
			w.WriteJson(&utils.Response{
				Status: "OK",
			})
		}
	}

	// Build the RESTful server
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// Control path upstream
		rest.Put("/$jade$/status", status),
		// Data path upstream
		rest.Post("/collector", receive),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	fmt.Println("video comb is listening on port", *port)
	log.Fatal(http.ListenAndServe(*addr+":"+strconv.Itoa(*port), api.MakeHandler()))
	fmt.Println("vido comb is done")
}
