package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var LIMIT = 5_000_000

func BenchmarkAPI(b *testing.B) {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "localhost:8080/fizzbuzz", nil)
	if err != nil {
		b.Fatalf("got error: %v", err)
	}

	req.URL.RawQuery = encodeQueryParams([]QueryParam{
		{"limit", strconv.Itoa(LIMIT)},
		{"int1", "3"},
		{"int2", "5"},
		{"str1", "fizz"},
		{"str2", "buzz"},
	})

	memCache := NewMemCache()
	memCache.FizzBuzz(recorder, req)

	assert.Equal(b, http.StatusOK, recorder.Code)
}
