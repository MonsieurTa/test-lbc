package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/MonsieurTa/fizzbuzz/pkg/fizzbuzz"
	"github.com/gorilla/schema"
)

type ParamsByHitCount map[string]int

func (memo ParamsByHitCount) MostUsed() (string, int) {
	maxKey, maxCount := "", 0
	for key, count := range memo {
		if count > maxCount {
			maxKey, maxCount = key, count
		}
	}
	return maxKey, maxCount
}

func (memo ParamsByHitCount) Empty() bool {
	return len(memo) == 0
}

type MemCache struct {
	mutex            sync.Mutex
	paramsByHitCount ParamsByHitCount
}

func NewMemCache() MemCache {
	return MemCache{paramsByHitCount: ParamsByHitCount{}}
}

func (c *MemCache) IncrementMostUsedRequest(key string) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	_, ok := c.paramsByHitCount[key]
	if !ok {
		c.paramsByHitCount[key] = 1
	} else {
		c.paramsByHitCount[key] += 1
	}
}

var decoder = schema.NewDecoder()

func (c *MemCache) FizzBuzz(w http.ResponseWriter, req *http.Request) {
	var cfg fizzbuzz.Config

	err := decoder.Decode(&cfg, req.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	service, err := fizzbuzz.New(cfg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	rv := service.Json()
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(rv))

	b, _ := json.Marshal(&cfg)
	c.IncrementMostUsedRequest(string(b))
}

type StatsResponse struct {
	Count  int             `json:"count"`
	Params fizzbuzz.Config `json:"params"`
}

func (c *MemCache) Stats(w http.ResponseWriter, req *http.Request) {
	var cfg fizzbuzz.Config

	if c.paramsByHitCount.Empty() {
		w.Write([]byte("No request yet"))
		return
	}

	params, count := c.paramsByHitCount.MostUsed()
	json.Unmarshal([]byte(params), &cfg)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StatsResponse{count, cfg})
}

func main() {
	withCache := NewMemCache()

	http.HandleFunc("/fizzbuzz", withCache.FizzBuzz)
	http.HandleFunc("/stats", withCache.Stats)

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
