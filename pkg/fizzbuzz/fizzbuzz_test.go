package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	service, err := New(Config{})

	assert.Nil(t, service)
	assert.Error(t, err)
	assert.Equal(t, "limit must be > 0", err.Error())

	service, err = New(Config{limit: 1, int1: -1})

	assert.Nil(t, service)
	assert.Error(t, err)
	assert.Equal(t, "input integers must be >= 0", err.Error())
}

func TestFizzBuzz(t *testing.T) {
	service, err := New(Config{limit: 1, int1: 1, int2: 1, fizz: "test", buzz: "ok"})

	assert.NotNil(t, service)
	assert.Nil(t, err)
}

func TestFizzBuzzGenerate(t *testing.T) {
	service, err := New(Config{limit: 15, int1: 3, int2: 5, fizz: "fizz", buzz: "buzz"})

	assert.NotNil(t, service)
	assert.Nil(t, err)

	actual := service.Generate()
	expected := []string{
		"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz",
	}
	assert.Equal(t, actual, expected)
}
