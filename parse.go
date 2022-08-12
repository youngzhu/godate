package godate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 日期转换

const defaultFormat = layoutDate0

func Parse(v interface{}) (Date, error) {

	switch v.(type) {
	case string:
		return ParseString(v.(string))
	default:
		return Date{}, fmt.Errorf("unsupported type: %T", v)
	}

	return NewDate(), nil
}

func ParseString(s string) (Date, error) {
	var (
		year  int
		month Month
		day   int

		validDate *regexp.Regexp

		err = fmt.Errorf("unrecognizable date: %s", s)
	)

	// 20210726
	if len(s) == 8 {
		year, month, day, err = parseStr3(s)
		if err == nil {
			return NewDateYMD(year, int(month), day)
		}
	}

	// 2021-07-26
	validDate = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if validDate.MatchString(s) {
		year, month, day, err = parseStr1(s)
		if err == nil {
			return NewDateYMD(year, int(month), day)
		}
	}

	// July 1st, 2021
	validDate = regexp.MustCompile(`[A-Z][a-z]+ \d{1,2}(st|nd|rd|th)?, \d{4}`)
	if validDate.MatchString(s) {
		year, month, day, err = parseStr2(s)
		if err == nil {
			return NewDateYMD(year, int(month), day)
		}
	}

	return Date{}, err
}

// layout: 20210726
func parseStr3(s string) (year int, month Month, day int, err error) {

	year, err = parseYear(s[0:4])
	if err != nil {
		return 0, 0, 0, err
	}

	month, err = parseMonth(s[4:6])
	if err != nil {
		return 0, 0, 0, err
	}

	day, err = parseDay(s[6:8])
	if err != nil {
		return 0, 0, 0, err
	}

	return year, month, day, nil
}

// layout: July 1st, 2021
func parseStr2(s string) (year int, month Month, day int, err error) {
	values := strings.Split(s, " ")
	//fmt.Println(values)

	month, err = parseMonth(values[0])
	if err != nil {
		return 0, 0, 0, err
	}

	day, err = parseDay(values[1])
	if err != nil {
		return 0, 0, 0, err
	}

	year, err = parseYear(values[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return year, month, day, nil
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
		return 0, fmt.Errorf("invalid day: %s", s)
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

	return 0, fmt.Errorf("invalid month: %s", mon)
}

// layout: 2021-07-26
func parseStr1(s string) (year int, month Month, day int, err error) {
	values := strings.Split(s, "-")
	//fmt.Println(values)
	y, err := strconv.ParseInt(values[0], 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid year: %s", values[0])
	}
	year = int(y)

	m, err := strconv.ParseInt(values[1], 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid month: %s", values[1])
	}
	month = Month(m)

	d, err := strconv.ParseInt(values[2], 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid year: %s", values[2])
	}
	day = int(d)

	return year, month, day, nil
}