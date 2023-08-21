package main

import (
	"fmt"
	"io"
	"net/http"
)

const portNumber = ":8080"

func ServerStart() {
	http.HandleFunc("/visitorstats", VisitorStatsHandler)
	http.ListenAndServe(portNumber, nil)
}

func VisitorStatsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GET_VisitorStats(w, r)
	case "POST":
		POST_IncrementVisitorStats(w, r)
	default:
		fmt.Fprintf(w, "Error: Invalid method")
	}

}

func GET_VisitorStats(w http.ResponseWriter, r *http.Request) {
	stats := DatabaseRead()
	w.Write(stats)
}

func POST_IncrementVisitorStats(w http.ResponseWriter, r *http.Request) {
	country_code, _ := io.ReadAll(r.Body)
	if len(country_code) != 2 {
		w.Write([]byte("Invalid country code! ! ! "))
	} else {
		reply := DatabaseWrite(country_code)
		w.Write(reply)
	}
}
