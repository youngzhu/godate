package godate

import "testing"

var parseTestCases = []struct {
	input string
	want  string
}{
	{"2021-07-26", "2021_07_26"},
	{"July 1st, 2021", "2021_07_01"},
}

func TestParseString(t *testing.T) {
	for _, tc := range parseTestCases {
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
