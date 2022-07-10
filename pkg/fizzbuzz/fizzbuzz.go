package fizzbuzz

import (
	"errors"
	"strconv"
)

type Config struct {
	Limit int    `schema:"limit"`
	Int1  int    `schema:"int1"`
	Int2  int    `schema:"int2"`
	Fizz  string `schema:"str1"`
	Buzz  string `schema:"str2"`
}

func (cfg *Config) valid() error {
	if cfg.Limit < 1 {
		return errors.New("limit must be > 0")
	}

	if cfg.Int1 == 0 {
		return errors.New("int1 can't 0")
	}

	if cfg.Int2 == 0 {
		return errors.New("int2 can't 0")
	}
	return nil
}

type Service struct {
	limit int
	int1  int
	int2  int
	fizz  string
	buzz  string
}

func New(cfg Config) (*Service, error) {
	err := cfg.valid()
	if err != nil {
		return nil, err
	}

	return &Service{
		limit: cfg.Limit,
		int1:  abs(cfg.Int1),
		int2:  abs(cfg.Int2),
		fizz:  cfg.Fizz,
		buzz:  cfg.Buzz,
	}, nil
}

func (s *Service) Generate() []string {
	rv := make([]string, s.limit)

	int1Multiple, int2Multiple := 0, 0
	for i := 1; i <= s.limit; i++ {
		int1Multiple++
		int2Multiple++

		if int1Multiple != s.int1 && int2Multiple != s.int2 {
			rv[i-1] += strconv.Itoa(i)
			continue
		}

		if int1Multiple == s.int1 {
			rv[i-1] += s.fizz
			int1Multiple = 0
		}
		if int2Multiple == s.int2 {
			rv[i-1] += s.buzz
			int2Multiple = 0
		}
	}
	return rv
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
