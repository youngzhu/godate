package godate

import "testing"

func TestMonth(t *testing.T) {
	t.Log(January)
}

func TestMonth_String(t *testing.T) {
	t.Log(January.String())
	t.Log(Month(13).String())
}

var monthDaysTests = []struct {
	year, month, days int
}{
	{2011, 1, 31},  // January, first month, 31 days
	{2011, 2, 28},  // February, non-leap year, 28 days
	{2012, 2, 29},  // February, leap year, 29 days
	{2011, 6, 30},  // June, 30 days
	{2011, 12, 31}, // December, last month, 31 days
}

func TestDate_MonthDays(t *testing.T) {
	for _, tt := range monthDaysTests {
		date, _ := NewDateYMD(tt.year, tt.month, 1)
		days := date.MonthDays()
		if days != tt.days {
			t.Errorf("got %d; expected %d for %d-%02d",
				days, tt.days, tt.year, tt.month)
		}
	}
}
