package godate

import (
	"fmt"
	"regexp"
	"testing"
)

var parseStringTestCases = []struct {
	input string
	want  string
}{
	{"20210726", "2021_07_26"},
	{"2021-07-26", "2021_07_26"},
	{"2021/07/26", "2021_07_26"},
	{"07/26/2021", "2021_07_26"},
	{"July 1st, 2021", "2021_07_01"},
}

func TestParseString(t *testing.T) {
	for _, tc := range parseStringTestCases {
		date, err := parseString0(tc.input)
		if err != nil {
			t.Error(err)
		}
		got := date.Format("2006_01_02")
		if got != tc.want {
			t.Errorf("got: %s; want: %s", got, tc.want)
		}
	}
}

func TestParse(t *testing.T) {
	testcases := []struct {
		input interface{}
		date  Date
	}{
		{"2023-10-01", MustDate(2023, 10, 1)},
		{"2023-9-1", MustDate(2023, 9, 1)},
		{"2023/10/01", MustDate(2023, 10, 1)},
		{"2023/9/1", MustDate(2023, 9, 1)},
		{"10/01/2023", MustDate(2023, 10, 1)},
		{"9/1/2023", MustDate(2023, 9, 1)},
		//{"20231001", MustDate(2023, 10, 1), nil},
		//{"2023.10.01", MustDate(2023, 10, 1), nil},
		//{"2023.10.1", MustDate(2023, 10, 1), nil},
		//{"2023.9.1", MustDate(2023, 9, 1), nil},
	}
	for i, tc := range testcases {
		name := fmt.Sprintf("No.%d", i+1)
		t.Run(name, func(t *testing.T) {
			date, err := Parse(tc.input)
			if err != nil {
				t.Errorf("got error: %v\n", err)
			}
			if err == nil {
				if !date.IsTheSameDay(tc.date) {
					t.Errorf("want: %v, but got: %v\n", tc.date, date)
				}
			}
		})
	}
}

// todo
//func TestParse_error(t *testing.T) {
//	testcases := []struct {
//		input interface{}
//		err   error
//	}{
//		{"2023-19-1", nil},
//	}
//	for i, tc := range testcases {
//		name := fmt.Sprintf("No.%d", i+1)
//		t.Run(name, func(t *testing.T) {
//			_, err := Parse(tc.input)
//			if !errors.Is(err, tc.err) {
//				t.Errorf("parse error, want: %v, but got: %v\n", tc.err, err)
//			}
//		})
//	}
//}

//
// test others
//
func TestRegExp_FindStringSubmatch(t *testing.T) {
	re := regexp.MustCompile(`^(\d{4})-(\d{1,2})$`)
	fmt.Printf("%q\n", re.FindStringSubmatch("2023-1"))
	fmt.Printf("%q\n", re.FindStringSubmatch("2023a"))
	fmt.Printf("%q\n", re.FindStringSubmatch("a2023"))

	fmt.Printf("%v\n", re.MatchString("2023"))
	fmt.Printf("%v\n", re.MatchString("2023a"))
	fmt.Printf("%v\n", re.MatchString("a2023"))

	fmt.Printf("%v\n", re.FindStringSubmatch("2023") == nil)
	fmt.Printf("%v\n", re.FindStringSubmatch("2023a") == nil)
	fmt.Printf("%v\n", re.FindStringSubmatch("a2023") == nil)
}
