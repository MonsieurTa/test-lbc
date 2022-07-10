package main

import (
	"encoding/json"
	"net/http"

	"github.com/MonsieurTa/fizzbuzz/pkg/fizzbuzz"
	"github.com/gorilla/schema"
)

type WithCache struct {
	requestsPerHit map[string]int
}

var decoder = schema.NewDecoder()

func (c *WithCache) FizzBuzz(w http.ResponseWriter, req *http.Request) {
	var cfg fizzbuzz.Config

	err := decoder.Decode(&cfg, req.URL.Query())
	if err != nil {
		w.WriteHeader(400)
		return
	}

	service, err := fizzbuzz.New(cfg)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	rv := service.Generate()
	err = json.NewEncoder(w).Encode(rv)
	if err != nil {
		w.WriteHeader(400)
	}
}

func (c *WithCache) Stats(w http.ResponseWriter, req *http.Request) {

}

func main() {
	withCache := WithCache{}

	http.HandleFunc("/fizzbuzz", withCache.FizzBuzz)
	http.HandleFunc("/stats", withCache.Stats)

	http.ListenAndServe(":8080", nil)
}
