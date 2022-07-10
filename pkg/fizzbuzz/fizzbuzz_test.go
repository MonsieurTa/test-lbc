package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		cfg      Config
		expected string
	}{
		{
			cfg:      Config{limit: 0},
			expected: "limit must be > 0",
		},
		{
			cfg:      Config{limit: -1},
			expected: "limit must be > 0",
		},
		{
			cfg:      Config{limit: 15, int1: 0, int2: 0, fizz: "fizz", buzz: "buzz"},
			expected: "int1 can't 0",
		},
		{
			cfg:      Config{limit: 15, int1: 1, int2: 0, fizz: "fizz", buzz: "buzz"},
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

func TestFizzBuzz(t *testing.T) {
	service, err := New(Config{limit: 1, int1: 1, int2: 1, fizz: "test", buzz: "ok"})

	assert.NotNil(t, service)
	assert.Nil(t, err)
}

func TestFizzBuzzGenerate(t *testing.T) {
	tests := []struct {
		cfg      Config
		expected []string
	}{
		{
			cfg:      Config{limit: 15, int1: 3, int2: 5, fizz: "fizz", buzz: "buzz"},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			cfg:      Config{limit: 10, int1: 2, int2: 2, fizz: "fizz", buzz: "buzz"},
			expected: []string{"1", "fizzbuzz", "3", "fizzbuzz", "5", "fizzbuzz", "7", "fizzbuzz", "9", "fizzbuzz"},
		},
		{
			cfg:      Config{limit: 10, int1: 1, int2: 2, fizz: "fizz", buzz: "buzz"},
			expected: []string{"fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz"},
		},
		{
			cfg:      Config{limit: 10, int1: 1, int2: 2, fizz: "", buzz: ""},
			expected: []string{"", "", "", "", "", "", "", "", "", ""},
		},
	}

	for _, test := range tests {
		service, err := New(test.cfg)

		assert.NotNil(t, service)
		assert.Nil(t, err)

		actual := service.Generate()
		assert.Equal(t, actual, test.expected)
	}
}
