package godate

import (
	"fmt"
	"regexp"
)

// 日期转换

var layouts = []string{
	"YYYY-MM-DD", // 2021-07-26
	"YYYY-MM-DD", // July 1st, 2021
}

func ParseString(s string) (Date, error) {
	var (
		year  int = -1
		month Month
		day   int
	)

	var validDate = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	if validDate.MatchString(s) {
		//s := validDate.FindAllString(s, -1)
		s := validDate.Split(`-`, -1)
		//s:=validDate.
		fmt.Println(s)
	}

	return NewDateYMD(year, int(month), day)
}

func Parse(v interface{}) (Date, error) {

	return NewDate(), nil
}
