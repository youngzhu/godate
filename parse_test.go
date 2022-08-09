package godate

import "testing"

var parseStringTestCases = []struct {
	input string
	want  string
}{
	{"2021-07-26", "2021_07_26"},
	{"July 1st, 2021", "2021_07_01"},
}

func TestParseString(t *testing.T) {
	for _, tc := range parseStringTestCases {
		date, err := ParseString(tc.input)
		if err != nil {
			t.Error(err)
		}
		got := date.Format("2006_01_02")
		if got != tc.want {
			t.Errorf("got: %s; want: %s", got, tc.want)
		}
	}
}

var parseTestCases = []struct {
	input interface{}
	want  string
}{
	{"2021-07-26", "2021-07-26"},
	//{0, "2021-07-26"},
}

func TestParse(t *testing.T) {
	for _, tc := range parseTestCases {
		date, err := Parse(tc.input)
		if err != nil {
			t.Error(err)
		}
		got := date.Format(defaultFormat)
		if got != tc.want {
			t.Errorf("got: %s; want: %s", got, tc.want)
		}
	}
}
