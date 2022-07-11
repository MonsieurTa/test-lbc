package fizzbuzz

import (
	"errors"
	"strconv"
)

type Config struct {
	Limit int32  `schema:"limit" json:"limit"`
	Int1  int32  `schema:"int1" json:"int1"`
	Int2  int32  `schema:"int2" json:"int2"`
	Str1  string `schema:"str1" json:"str1"`
	Str2  string `schema:"str2" json:"str2"`
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
	limit int32
	int1  int32
	int2  int32
	str1  string
	str2  string
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
		str1:  cfg.Str1,
		str2:  cfg.Str2,
	}, nil
}

func (s *Service) Generate() []string {
	rv := make([]string, s.limit)

	int1Multiple, int2Multiple := int32(0), int32(0)
	for i := 1; i <= int(s.limit); i++ {
		int1Multiple++
		int2Multiple++

		if int1Multiple != s.int1 && int2Multiple != s.int2 {
			rv[i-1] += strconv.Itoa(i)
			continue
		}

		if int1Multiple == s.int1 {
			rv[i-1] += s.str1
			int1Multiple = 0
		}
		if int2Multiple == s.int2 {
			rv[i-1] += s.str2
			int2Multiple = 0
		}
	}
	return rv
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
