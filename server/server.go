package server

import (
	"log"
	"net/http"
)

func Execute() {
	hh := HttpHandlers{}
	http.HandleFunc("/add_request", hh.AddReqFunc)
	http.HandleFunc("/get_stats", hh.GetStatsFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
