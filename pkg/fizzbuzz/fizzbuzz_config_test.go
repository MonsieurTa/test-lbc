package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Returns_Expected_Error(t *testing.T) {
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
