package godate

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 日期转换

var (
	ErrIllegalArgument = errors.New("illegal argument")
	ErrInvalidMonth    = errors.New("invalid month value")
	ErrInvalidDay      = errors.New("invalid day value")
	ErrUnrecognizable  = errors.New("unrecognizable date")
)

type parseError struct {
	cause   error  // the underlying error
	message string // describes the error
}

func (p *parseError) Error() string {
	return fmt.Sprintf("parse error, %s: %s", p.cause, p.message)
}

func (p *parseError) Unwrap() error {
	return p.cause
}

func newParseError(cause error, message string) *parseError {
	return &parseError{
		cause:   cause,
		message: message,
	}
}

func Parse(v interface{}) (Date, error) {
	switch v.(type) {
	case string:
		return parseString(v.(string))
	default:
		errMsg := fmt.Sprintf("unsupported type: %T", v)
		return Date{}, newParseError(ErrIllegalArgument, errMsg)
	}
}

func parseString(s string) (Date, error) {
	var exp *regexp.Regexp
	var matches []string

	if strings.ContainsRune(s, '-') {
		//yyyy-mm-dd
		exp = regexp.MustCompile(`^(\d{4})-(\d{1,2})-(\d{1,2})$`)
		matches = exp.FindStringSubmatch(s)
		if matches != nil {
			return parseYMDSlice(matches[1:])
		}
	} else if strings.ContainsRune(s, '/') {
		// yyyy/mm/dd
		exp = regexp.MustCompile(`^(\d{4})/(\d{1,2})/(\d{1,2})$`)
		matches = exp.FindStringSubmatch(s)
		if matches != nil {
			return parseYMDSlice(matches[1:])
		} else {
			// mm/dd/yyyy
			exp = regexp.MustCompile(`^(\d{1,2})/(\d{1,2})/(\d{4})$`)
			matches = exp.FindStringSubmatch(s)
			if matches != nil {
				ymd := make([]string, 3)
				ymd[0] = matches[3]
				ymd[1] = matches[1]
				ymd[2] = matches[2]
				return parseYMDSlice(ymd)
			}
		}
	} else if strings.ContainsRune(s, '.') {
		//yyyy.mm.dd
		exp = regexp.MustCompile(`^(\d{4}).(\d{1,2}).(\d{1,2})$`)
		matches = exp.FindStringSubmatch(s)
		if matches != nil {
			return parseYMDSlice(matches[1:])
		}
	} else if len(s) == 8 {
		//yyyymmdd
		exp = regexp.MustCompile(`^(\d{4})(\d{1,2})(\d{1,2})$`)
		matches = exp.FindStringSubmatch(s)
		if matches != nil {
			return parseYMDSlice(matches[1:])
		}
	}

	// July 1st, 2021
	exp = regexp.MustCompile(`^([A-Z][a-z]+) (\d{1,2})(st|nd|rd|th)?, (\d{4})$`)
	matches = exp.FindStringSubmatch(s)
	if matches != nil {
		var year, month, day string
		if len(matches) == 4 {
			year, month, day = matches[3], matches[1], matches[2]
		} else {
			year, month, day = matches[4], matches[1], matches[2]
		}

		return parseYMDSlice([]string{year, month, day})
	}

	return Date{}, newParseError(ErrUnrecognizable, s)
}

func parseYMDSlice(ymd []string) (date Date, err error) {
	var year, day int
	var month Month

	year, err = parseYear(ymd[0])
	if err != nil {
		goto fail
	}
	month, err = parseMonth(ymd[1])
	if err != nil {
		goto fail
	}
	day, err = parseDay(ymd[2])
	if err != nil {
		goto fail
	}

	date = MustDate(year, month.IntValue(), day)

fail:
	return
}

func parseYear(s string) (int, error) {
	y, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(y), nil
}

func parseDay(s string) (int, error) {
	d := s
	if len(s) >= 3 {
		d = s[0 : len(s)-3]
	}

	dd, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		return 0, err
	}
	dayInt := int(dd)
	if 1 <= dayInt && dayInt <= 31 {
		return dayInt, nil
	} else {
		return 0, newParseError(ErrInvalidDay, s)
	}
}

func parseMonth(mon string) (Month, error) {
	if len(mon) == 3 {
		for i, v := range shortMonthNames {
			if v == mon {
				return Month(i + 1), nil
			}
		}
	} else {
		num := regexp.MustCompile(`\d{1,2}`)
		if num.MatchString(mon) {
			m, err := strconv.ParseInt(mon, 10, 32)
			if err == nil {
				mm := Month(m)
				if January <= mm && mm <= December {
					return mm, nil
				}
			}

		} else {
			for i, v := range longMonthNames {
				if v == mon {
					return Month(i + 1), nil
				}
			}
		}
	}

	return 0, newParseError(ErrInvalidMonth, mon)
}
