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

	return &Service{cfg: &cfg}, nil
}

func (s *Service) Generate() []string {
	var v string
	rv := make([]string, s.cfg.limit)

	fizzbuzz := s.cfg.fizz + s.cfg.buzz
	for i := 1; i <= s.cfg.limit; i++ {
		if i%s.cfg.int1 == 0 && i%s.cfg.int2 == 0 {
			v = fizzbuzz
		} else if i%s.cfg.int1 == 0 {
			v = s.cfg.fizz
		} else if i%s.cfg.int2 == 0 {
			v = s.cfg.buzz
		} else {
			v = strconv.Itoa(i)
		}
		rv[i-1] = v
	}
	return rv
}
