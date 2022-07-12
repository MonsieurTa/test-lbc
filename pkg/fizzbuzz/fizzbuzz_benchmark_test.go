package fizzbuzz

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var LIMIT = 5_000_000

func BenchmarkFizzBuzz_Array_With_High_Limit(b *testing.B) {
	service, err := New(Config{
		Limit: LIMIT,
		Int1:  3,
		Int2:  5,
		Str1:  "fizz",
		Str2:  "buzz",
	})

	assert.NotNil(b, service)
	assert.Nil(b, err)

	json.Marshal(service.Array())
}

func BenchmarkFizzBuzz_Json_With_High_Limit(b *testing.B) {
	service, err := New(Config{
		Limit: LIMIT,
		Int1:  3,
		Int2:  5,
		Str1:  "fizz",
		Str2:  "buzz",
	})

	assert.NotNil(b, service)
	assert.Nil(b, err)

	// already JSON formatted
	service.Json()
}
