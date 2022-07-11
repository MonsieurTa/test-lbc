package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type QueryParam [2]string

func encodeQueryParams(queryParams []QueryParam) string {
	query := url.Values{}
	for _, params := range queryParams {
		key, value := params[0], params[1]
		query.Add(key, value)
	}
	return query.Encode()
}

func TestFizzBuzzHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "localhost:8080/fizzbuzz", nil)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	req.URL.RawQuery = encodeQueryParams([]QueryParam{
		{"limit", "15"},
		{"int1", "3"},
		{"int2", "5"},
		{"str1", "fizz"},
		{"str2", "buzz"},
	})

	withCache := NewCache()
	withCache.FizzBuzz(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var actual []string
	_ = json.Unmarshal(recorder.Body.Bytes(), &actual)

	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	assert.Equal(t, expected, actual)
}

func TestStats(t *testing.T) {
	fizzBuzzRecorder := httptest.NewRecorder()
	fizzBuzzReq, err := http.NewRequest("GET", "localhost:8080/fizzbuzz", nil)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	fizzBuzzReq.URL.RawQuery = encodeQueryParams([]QueryParam{
		{"limit", "15"},
		{"int1", "3"},
		{"int2", "5"},
		{"str1", "fizz"},
		{"str2", "buzz"},
	})

	withCache := NewCache()

	for i := 0; i < 1000; i++ {
		withCache.FizzBuzz(fizzBuzzRecorder, fizzBuzzReq)
		assert.Equal(t, http.StatusOK, fizzBuzzRecorder.Code)
	}

	statsRecorder := httptest.NewRecorder()
	statsReq, err := http.NewRequest("GET", "localhost:8080/stats", nil)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	withCache.Stats(statsRecorder, statsReq)
	fmt.Println(statsRecorder)
}
