package main

import (
	"adventofcode2024/handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	httpAddr = flag.String("addr", "0.0.0.0:8080", "Listen address")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", handlers.MainHandler)

	for day := 1; day <= 25; day++ {
		http.HandleFunc(fmt.Sprintf("/day%d", day), handlers.DayHandler(day))
	}

	log.Printf("Serving http://%s", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))

}
