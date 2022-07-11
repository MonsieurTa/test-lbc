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

type Cache struct {
	mutex            sync.Mutex
	paramsByHitCount ParamsByHitCount
}

func NewCache() Cache {
	return Cache{paramsByHitCount: ParamsByHitCount{}}
}

func (c *Cache) IncrementMostUsedRequest(key string) {
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

func (c *Cache) FizzBuzz(w http.ResponseWriter, req *http.Request) {
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

	rv := service.Generate()
	err = json.NewEncoder(w).Encode(rv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	b, _ := json.Marshal(&cfg)
	c.IncrementMostUsedRequest(string(b))
}

func (c *Cache) Stats(w http.ResponseWriter, req *http.Request) {
	var cfg fizzbuzz.Config

	params, count := c.paramsByHitCount.MostUsed()
	json.Unmarshal([]byte(params), &cfg)

	json.NewEncoder(w).Encode(struct {
		Count  int             `json:"count"`
		Params fizzbuzz.Config `json:"params"`
	}{count, cfg})
}

func main() {
	withCache := NewCache()

	http.HandleFunc("/fizzbuzz", withCache.FizzBuzz)
	http.HandleFunc("/stats", withCache.Stats)

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
