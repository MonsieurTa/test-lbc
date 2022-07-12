package fizzbuzz

import (
	"errors"
	"strconv"
	"strings"
)

type Config struct {
	Limit int    `schema:"limit" json:"limit"`
	Int1  int    `schema:"int1" json:"int1"`
	Int2  int    `schema:"int2" json:"int2"`
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
	limit int
	int1  int
	int2  int
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

func (s *Service) String() string {
	var builder strings.Builder

	builder.WriteRune('[')
	int1Multiple, int2Multiple := 0, 0
	for i := 1; i <= s.limit; i++ {
		int1Multiple++
		int2Multiple++

		if i == 1 {
			builder.WriteString(`"`)
		} else {
			builder.WriteString(`,"`)
		}

		if int1Multiple != s.int1 && int2Multiple != s.int2 {
			builder.WriteString(strconv.FormatInt(int64(i), 10))
		} else {
			if int1Multiple == s.int1 {
				builder.WriteString(s.str1)
				int1Multiple = 0
			}
			if int2Multiple == s.int2 {
				builder.WriteString(s.str2)
				int2Multiple = 0
			}
		}
		builder.WriteString(`"`)
	}
	builder.WriteRune(']')
	return builder.String()
}

func (s *Service) Array() []string {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
