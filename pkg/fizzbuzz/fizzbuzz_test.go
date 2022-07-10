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
			cfg:      Config{Limit: 0},
			expected: "limit must be > 0",
		},
		{
			cfg:      Config{Limit: -1},
			expected: "limit must be > 0",
		},
		{
			cfg: Config{
				Limit: 15,
				Int1:  0,
				Int2:  0,
				Fizz:  "fizz",
				Buzz:  "buzz",
			},
			expected: "int1 can't 0",
		},
		{
			cfg: Config{
				Limit: 15,
				Int1:  1,
				Int2:  0,
				Fizz:  "fizz",
				Buzz:  "buzz",
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

func TestFizzBuzz(t *testing.T) {
	service, err := New(Config{Limit: 1, Int1: 1, Int2: 1, Fizz: "test", Buzz: "ok"})

	assert.NotNil(t, service)
	assert.Nil(t, err)
}

func TestFizzBuzzGenerate(t *testing.T) {
	tests := []struct {
		cfg      Config
		expected []string
	}{
		{
			cfg: Config{
				Limit: 15,
				Int1:  3,
				Int2:  5,
				Fizz:  "fizz",
				Buzz:  "buzz",
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			cfg: Config{
				Limit: 10,
				Int1:  2,
				Int2:  2,
				Fizz:  "fizz",
				Buzz:  "buzz",
			},
			expected: []string{"1", "fizzbuzz", "3", "fizzbuzz", "5", "fizzbuzz", "7", "fizzbuzz", "9", "fizzbuzz"},
		},
		{
			cfg: Config{
				Limit: 10,
				Int1:  1,
				Int2:  2,
				Fizz:  "fizz",
				Buzz:  "buzz",
			},
			expected: []string{"fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz"},
		},
		{
			cfg: Config{
				Limit: 10,
				Int1:  1,
				Int2:  2,
				Fizz:  "",
				Buzz:  "",
			},
			expected: []string{"", "", "", "", "", "", "", "", "", ""},
		},
		{
			cfg: Config{
				Limit: 10,
				Int1:  20,
				Int2:  20,
				Fizz:  "",
				Buzz:  "",
			},
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
		{
			cfg: Config{
				Limit: 10,
				Int1:  -1,
				Int2:  -2,
				Fizz:  "fizz",
				Buzz:  "buzz",
			},
			expected: []string{"fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz", "fizzbuzz"},
		},
	}

	for _, test := range tests {
		service, err := New(test.cfg)

		assert.NotNil(t, service)
		assert.Nil(t, err)

		actual := service.Generate()
		assert.Equal(t, test.expected, actual)
	}
}
