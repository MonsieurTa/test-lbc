package fizzbuzz

import (
	"errors"
	"strconv"
)

type Config struct {
	limit int

	int1 int
	int2 int

	fizz string
	buzz string
}

func (c *Config) valid() error {
	if c.limit <= 0 {
		return errors.New("limit must be > 0")
	}

	if c.int1 == 0 {
		return errors.New("int1 can't 0")
	}

	if c.int2 == 0 {
		return errors.New("int2 can't 0")
	}
	return nil
}

type Service struct {
	cfg *Config
}

func New(cfg Config) (*Service, error) {
	err := cfg.valid()
	if err != nil {
		return nil, err
	}

	cfg.int1 = abs(cfg.int1)
	cfg.int2 = abs(cfg.int2)

	return &Service{cfg: &cfg}, nil
}

func (s *Service) Generate() []string {
	rv := make([]string, s.cfg.limit)

	int1Multiple, int2Multiple := 0, 0
	for i := 1; i <= s.cfg.limit; i++ {
		int1Multiple += 1
		int2Multiple += 1

		if int1Multiple != s.cfg.int1 && int2Multiple != s.cfg.int2 {
			rv[i-1] += strconv.Itoa(i)
			continue
		}

		if int1Multiple == s.cfg.int1 {
			rv[i-1] += s.cfg.fizz
			int1Multiple = 0
		}
		if int2Multiple == s.cfg.int2 {
			rv[i-1] += s.cfg.buzz
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
