package meander

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var APIKey string

type Place struct {
*googleGeometry `json:"geometry"`
Name string `json:"name"`
Icon string `json:"icon"`
Photos []*googlePhoto `json:"photos"`
Vicinity string `json:"vicinity"`
}
type googleResponse struct {
Results []*Place `json:"results"`
}
type googleGeometry struct {
*googleLocation `json:"location"`
}
type googleLocation struct {
Lat float64 `json:"lat"`
Lng float64 `json:"lng"`
}
type googlePhoto struct {
PhotoRef string `json:"photo_reference"`
URL string `json:"url"`
}

type Query struct {
	Lat float64
	Lng float64
	Journey []string
	Radius int
	CostRangeStr string
	}

func (p *Place) Public() interface{} {
	return map[string]interface{}{
	"name": p.Name,
	"icon": p.Icon,
	"photos": p.Photos,
	"vicinity": p.Vicinity,
	"lat": p.Lat,
	"lng": p.Lng,
	}
	}
	//func to call api
	func (q *Query) find(types string) (*googleResponse, error) {
		u :=
		"https://api.geoapify.com/v2/places"
		vals := make(url.Values)
		vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
		vals.Set("radius", fmt.Sprintf("%d", q.Radius))
		vals.Set("types", types)
		vals.Set("key", APIKey)
		if len(q.CostRangeStr) > 0 {
		r := ParseCostRange(q.CostRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
		}
		res, err := http.Get(u + "?" + vals.Encode())
if err != nil {
return nil, err
}
defer res.Body.Close()
var response googleResponse
if err := json.NewDecoder(res.Body).Decode(&response); err !=
nil {
return nil, err
}
return &response, nil
}

// Run runs the query concurrently, and returns the results.
func (q *Query) Run() []interface{} {
	rand.Seed(time.Now().UnixNano())
	var w sync.WaitGroup
	var l sync.Mutex
	places := make([]interface{}, len(q.Journey))
	for i, r := range q.Journey {
		w.Add(1)
		go func(types string, i int) {
		defer w.Done()
		response, err := q.find(types)
	if err != nil {
		log.Println("Failed to find places:", err)
		return
		}
		if len(response.Results) == 0 {
		log.Println("No places found for", types)
		return
		}
		for _, result := range response.Results {
		for _, photo := range result.Photos {
			photo.URL = photo.URL 
		}
		}
		randI := rand.Intn(len(response.Results))
		l.Lock()
		places[i] = response.Results[randI]
		l.Unlock()
		}(r, i)
		}
		w.Wait() // wait for everything to finish
		return places
		}