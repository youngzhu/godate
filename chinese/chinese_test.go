package chinese

import (
	"github.com/youngzhu/godate"
	"testing"
)

func TestIsOffDayInChina(t *testing.T) {
	testCases := []struct {
		date godate.Date
		want bool
	}{
		{godate.MustDate(2024, 1, 1), true},
		{godate.MustDate(2024, 1, 24), false},
		{godate.MustDate(2024, 2, 9), false},
		{godate.MustDate(2024, 2, 10), true},
		{godate.MustDate(2024, 2, 11), true},
		{godate.MustDate(2024, 2, 12), true},
		{godate.MustDate(2024, 2, 13), true},
		{godate.MustDate(2024, 2, 14), true},
		{godate.MustDate(2024, 2, 15), true},
		{godate.MustDate(2024, 2, 16), true},
		{godate.MustDate(2024, 2, 17), true},
		{godate.MustDate(2024, 2, 18), false},
		{godate.MustDate(2025, 1, 1), true},
		{godate.MustDate(2025, 10, 1), false},
	}

	for _, testcase := range testCases {
		t.Run("", func(t *testing.T) {
			got := IsOffDayInChina(testcase.date)
			if got != testcase.want {
				t.Errorf("%v is off day in China? want: %v, but got: %v",
					testcase.date, testcase.want, got)
			}
		})
	}
}

func TestIsWorkDayInChina(t *testing.T) {
	testCases := []struct {
		date godate.Date
		want bool
	}{
		{godate.MustDate(2024, 1, 1), false},
		{godate.MustDate(2024, 2, 4), true},
		{godate.MustDate(2024, 2, 9), true},
		{godate.MustDate(2024, 2, 17), false},
		{godate.MustDate(2024, 2, 18), true},
		{godate.MustDate(2026, 1, 1), false},
		{godate.MustDate(2026, 1, 3), false},
		{godate.MustDate(2026, 1, 4), true},
	}

	for _, testcase := range testCases {
		t.Run("", func(t *testing.T) {
			got := IsWorkDayInChina(testcase.date)
			if got != testcase.want {
				t.Errorf("%v is work day in China? want: %v, but got: %v",
					testcase.date, testcase.want, got)
			}
		})
	}
}

// func TestTimorFetcher_Fetch_fail(t *testing.T) {
// 	GetHolidays(2023)
// }
