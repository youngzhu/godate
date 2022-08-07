package godate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 日期转换

var layouts = []string{
	"YYYY-MM-DD", // 2021-07-26
	"YYYY-MM-DD", // July 1st, 2021
}

func ParseString(s string) (Date, error) {
	var (
		year  int
		month Month
		day   int
	)

	var validDate = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if validDate.MatchString(s) {
		values := strings.Split(s, "-")
		//fmt.Println(values)
		y, err := strconv.ParseInt(values[0], 10, 32)
		if err != nil {
			return Date{}, fmt.Errorf("invalid year: %s", values[0])
		}
		year = int(y)

		m, err := strconv.ParseInt(values[1], 10, 32)
		if err != nil {
			return Date{}, fmt.Errorf("invalid month: %s", values[1])
		}
		month = Month(m)

		d, err := strconv.ParseInt(values[2], 10, 32)
		if err != nil {
			return Date{}, fmt.Errorf("invalid year: %s", values[2])
		}
		day = int(d)
	}

	return NewDateYMD(year, int(month), day)
}

func Parse(v interface{}) (Date, error) {

	return NewDate(), nil
}
