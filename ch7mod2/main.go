package main

import (
	"ch7mod2/meander"
	"encoding/json"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	meander.APIKey = "YOUR_API_KEY"
	http.HandleFunc("/journeys", cors(func(w http.ResponseWriter, r *http.Request) {
	respond(w, r, meander.Journeys)
	}))
	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
	q := &meander.Query{
	Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
	}
	q.Lat, _ = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	q.Lng, _ = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	q.Radius, _ = strconv.Atoi(r.URL.Query().Get("radius"))
	q.CostRangeStr = r.URL.Query().Get("cost")
	places := q.Run()
	respond(w, r, places)
	}))
	http.ListenAndServe(":6521", http.DefaultServeMux)
}


func respond(w http.ResponseWriter, r *http.Request, data [] interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
	publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	f(w, r)
	}
	} 