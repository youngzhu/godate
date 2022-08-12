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

	ymd := s

	switch len(s) {
	case 8:
		// 20210726
		ymd = s
		fallthrough
	case 10:
		if strings.ContainsRune(s, '-') {
			// 2021-07-26
			ymd = strings.ReplaceAll(s, "-", "")
		} else if strings.ContainsRune(s, '/') {
			if strings.Index(s, `/`) == 4 {
				// 2021/07/26
				ymd = strings.ReplaceAll(s, `/`, "")
			} else {
				// 07/26/2021
				arr := strings.Split(s, `/`)
				ymd = arr[2] + arr[0] + arr[1]
			}
		}

		year, month, day, err = parseYMD(ymd)
		if err == nil {
			return NewDateYMD(year, int(month), day)
		}
	}

	// July 1st, 2021
	validDate = regexp.MustCompile(`[A-Z][a-z]+ \d{1,2}(st|nd|rd|th)?, \d{4}`)
	if validDate.MatchString(s) {
		year, month, day, err = parseLongDate(s)
		if err == nil {
			return NewDateYMD(year, int(month), day)
		}
	}

	return Date{}, err
}

// layout: 20210726
func parseYMD(s string) (year int, month Month, day int, err error) {

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
func parseLongDate(s string) (year int, month Month, day int, err error) {
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
