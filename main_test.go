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

func TestAPI_Returns_A_Valid_JSON(t *testing.T) {
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

	memCache := NewMemCache()
	memCache.FizzBuzz(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var actual []string
	_ = json.Unmarshal(recorder.Body.Bytes(), &actual)

	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	assert.Equal(t, expected, actual)
}

func TestAPI_Returns_Most_Used_Request(t *testing.T) {
	fizzBuzzRecorder := httptest.NewRecorder()
	fizzBuzzReq, err := http.NewRequest("GET", "localhost:8080/fizzbuzz", nil)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	memCache := NewMemCache()

	fizzBuzzReq.URL.RawQuery = encodeQueryParams([]QueryParam{
		{"limit", "15"},
		{"int1", "3"},
		{"int2", "5"},
		{"str1", "fizz"},
		{"str2", "buzz"},
	})

	for i := 0; i < 1000; i++ {
		memCache.FizzBuzz(fizzBuzzRecorder, fizzBuzzReq)
		assert.Equal(t, http.StatusOK, fizzBuzzRecorder.Code)
	}

	for i := 0; i < 100; i++ {
		fizzBuzzReq.URL.RawQuery = encodeQueryParams([]QueryParam{
			{"limit", "15"},
			{"int1", "3"},
			{"int2", "5"},
			{"str1", fmt.Sprintf("fizz#%d", i)},
			{"str2", fmt.Sprintf("buzz#%d", i)},
		})
		memCache.FizzBuzz(fizzBuzzRecorder, fizzBuzzReq)
		assert.Equal(t, http.StatusOK, fizzBuzzRecorder.Code)
	}

	statsRecorder := httptest.NewRecorder()
	statsReq, err := http.NewRequest("GET", "localhost:8080/stats", nil)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	memCache.Stats(statsRecorder, statsReq)

	var response StatsResponse
	json.Unmarshal(statsRecorder.Body.Bytes(), &response)

	assert.Equal(t, 1000, response.Count)
	assert.Equal(t, 15, response.Params.Limit)
	assert.Equal(t, 3, response.Params.Int1)
	assert.Equal(t, 5, response.Params.Int2)
	assert.Equal(t, "fizz", response.Params.Str1)
	assert.Equal(t, "buzz", response.Params.Str2)
}
