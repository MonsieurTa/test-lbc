package fizzbuzz

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Returns_Expected_Error(t *testing.T) {
	tests := []struct {
		cfg      Config
		expected string
	}{
		{
			cfg:      Config{},
			expected: "limit must be greater than 0",
		},
		{
			cfg:      Config{Limit: 0},
			expected: "limit must be greater than 0",
		},
		{
			cfg:      Config{Limit: -1},
			expected: "limit must be greater than 0",
		},
		{
			cfg: Config{
				Limit: 15,
				Int1:  0,
				Int2:  0,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expected: "int1 can't 0",
		},
		{
			cfg: Config{
				Limit: 15,
				Int1:  1,
				Int2:  0,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expected: "int2 can't 0",
		},
	}

	for _, test := range tests {
		service, err := New(test.cfg)

		assert.Nil(t, service)
		assert.Error(t, err)
		assert.Equal(t, test.expected, err.Error())
	}
}

var Tests = []struct {
	cfg      Config
	expected []string
}{
	{
		cfg: Config{
			Limit: 15,
			Int1:  3,
			Int2:  5,
			Str1:  "fizz",
			Str2:  "buzz",
		},
		expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
	},
	{
		cfg: Config{
			Limit: 10,
			Int1:  2,
			Int2:  2,
			Str1:  "fizz",
			Str2:  "buzz",
		},
		expected: []string{"1", "fizzbuzz", "3", "fizzbuzz", "5", "fizzbuzz", "7", "fizzbuzz", "9", "fizzbuzz"},
	},
	{
		cfg: Config{
			Limit: 10,
			Int1:  1,
			Int2:  2,
			Str1:  "fizz",
			Str2:  "buzz",
		},
		expected: []string{"fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz"},
	},
	{
		cfg: Config{
			Limit: 10,
			Int1:  1,
			Int2:  2,
			Str1:  "",
			Str2:  "",
		},
		expected: []string{"", "", "", "", "", "", "", "", "", ""},
	},
	{
		cfg: Config{
			Limit: 10,
			Int1:  20,
			Int2:  20,
			Str1:  "",
			Str2:  "",
		},
		expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		cfg: Config{
			Limit: 10,
			Int1:  -1,
			Int2:  -2,
			Str1:  "fizz",
			Str2:  "buzz",
		},
		expected: []string{"fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz"},
	},
}

func TestFizzBuzz_Array_With_Valid_Inputs(t *testing.T) {
	for _, test := range Tests {
		service, err := New(test.cfg)

		assert.NotNil(t, service)
		assert.Nil(t, err)

		actual := service.Array()
		assert.Equal(t, test.expected, actual)
	}
}

func TestFizzBuzz_String_With_Valid_Inputs(t *testing.T) {
	for _, test := range Tests {
		service, err := New(test.cfg)

		assert.NotNil(t, service)
		assert.Nil(t, err)

		actual := service.Json()

		expected, err := json.Marshal(test.expected)
		if err != nil {
			t.Fatalf("unexpected test error: %v", err)
		}
		assert.Equal(t, string(expected), actual)
	}
}
