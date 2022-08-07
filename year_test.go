package godate

import "testing"

var yearDaysTests = []struct {
	year, days int
}{
	{2011, 365}, // non-leap year, 355 days
	{2012, 366}, // leap year, 356 days
	{2020, 366}, // leap year, 356 days
	{2022, 365}, // non-leap year, 355 days
}

func TestDate_YearDays(t *testing.T) {
	for _, tt := range yearDaysTests {
		date, _ := NewDateYMD(tt.year, 1, 1)
		days := date.YearDays()
		if days != tt.days {
			t.Errorf("got %d; expected %d for year %d",
				days, tt.days, tt.year)
		}
	}
}
