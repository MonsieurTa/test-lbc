package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarFizzBuzz_Array_With_High_Limit(b *testing.B) {
	service, err := New(Config{
		Limit: 100_000_000,
		Int1:  3,
		Int2:  5,
		Str1:  "fizz",
		Str2:  "buzz",
	})

	assert.NotNil(b, service)
	assert.Nil(b, err)

	service.Array()
}

func BenchmarFizzBuzz_String_With_High_Limit(b *testing.B) {
	service, err := New(Config{
		Limit: 100_000_000,
		Int1:  3,
		Int2:  5,
		Str1:  "fizz",
		Str2:  "buzz",
	})

	assert.NotNil(b, service)
	assert.Nil(b, err)

	service.String()
}
