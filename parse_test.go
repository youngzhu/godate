package godate

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestParse_string(t *testing.T) {
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
		{"2023.10.01", MustDate(2023, 10, 1)},
		{"2023.10.1", MustDate(2023, 10, 1)},
		{"2023.9.1", MustDate(2023, 9, 1)},
		{"20231001", MustDate(2023, 10, 1)},
		{"20230901", MustDate(2023, 9, 1)},
		{"20230931", MustDate(2023, 9, 31)},
		{"20240101", MustDate(2024, 1, 1)},
		{"July 1st, 2021", MustDate(2021, 7, 1)},
		{"July 1, 2021", MustDate(2021, 7, 1)},
		{"July 11, 2021", MustDate(2021, 7, 11)},
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

func TestParse_time(t *testing.T) {
	testcases := []struct {
		input interface{}
		date  Date
	}{
		{time.Now(), today()},
		{time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local), MustDate(2023, 10, 1)},
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

func TestParse_error(t *testing.T) {
	testcases := []struct {
		input interface{}
		err   error
	}{
		{1.0, ErrIllegalArgument},
		{"2023-19-1", ErrInvalidMonth},
		{"2023-1-32", ErrInvalidDay},
		{"20231.3", ErrUnrecognizable},
	}
	for i, tc := range testcases {
		name := fmt.Sprintf("No.%d", i+1)
		t.Run(name, func(t *testing.T) {
			_, err := Parse(tc.input)
			if !errors.Is(err, tc.err) {
				t.Errorf("Expected error %q, got %q instead", tc.err, err)
			}
		})
	}
}

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
