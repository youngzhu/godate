package godate

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	t.Log(NewDate())
}

func TestNewDateYMD(t *testing.T) {
	year, month, day := 2000, 1, 1
	d, _ := NewDateYMD(year, month, day)
	if d.Year() != year || d.Month().IntValue() != month || d.Day() != day {
		t.Errorf("  want=year:%d, month:%d, day:%d", year, month, day)
		t.Errorf("  have=year:%d, month:%d, day:%d", d.Year(), d.Month(), d.Day())
	}
}

var addDayTests = []struct {
	date   Date
	days   int
	result Date
}{
	{MustDate(2022, 6, 18), 0, MustDate(2022, 6, 18)},
	{MustDate(2022, 6, 18), 1, MustDate(2022, 6, 19)},
	{MustDate(2022, 6, 18), 2, MustDate(2022, 6, 20)},
}

func TestDate_AddDay(t *testing.T) {
	for _, tt := range addDayTests {
		date, _ := tt.date.AddDay(tt.days)
		if !date.IsTheSameDay(tt.result) {
			t.Errorf("got %s; want %s", date, tt.result)
		}
	}
}

var subDayTests = []struct {
	date   Date
	days   int
	result Date
}{
	{MustDate(2022, 6, 18), 0, MustDate(2022, 6, 18)},
	{MustDate(2022, 6, 18), 1, MustDate(2022, 6, 17)},
	{MustDate(2022, 6, 18), 2, MustDate(2022, 6, 16)},
}

func TestDate_SubDay(t *testing.T) {
	for _, tt := range subDayTests {
		date, _ := tt.date.SubDay(tt.days)
		if !date.IsTheSameDay(tt.result) {
			t.Errorf("got %s; want %s", date, tt.result)
		}
	}
}

var beforeAfter = []struct {
	d1     Date
	d2     Date
	before bool
	after  bool
}{
	{Today(), Today(), false, false},
	{Today(), Tomorrow(), true, false},
	{Today(), Yesterday(), false, true},
	{MustDate(2023, 1, 1), MustDate(2023, 1, 2), true, false},
	{MustDate(2023, 1, 1), MustDate(2022, 12, 31), false, true},
	{MustDate(2023, 10, 1), MustDate(2023, 10, 1), false, false},
	{
		d1:     Date{time.Date(2023, time.September, 26, 20, 59, 59, 0, time.UTC)},
		d2:     Date{time.Date(2023, time.September, 26, 20, 59, 59, 0, time.UTC)},
		before: false,
		after:  false,
	},
	{
		Date{time.Date(2023, time.September, 26, 20, 59, 59, 0, time.UTC)},
		Date{time.Date(2023, time.September, 26, 21, 0, 0, 0, time.UTC)},
		true,
		false,
	},
	{
		Date{time.Date(2023, time.September, 26, 20, 59, 59, 0, time.UTC)},
		Date{time.Date(2023, time.September, 26, 20, 59, 58, 0, time.UTC)},
		false,
		true,
	},
}

func TestDate_Before(t *testing.T) {
	for _, testcase := range beforeAfter {
		t.Run("", func(t *testing.T) {
			got := testcase.d1.Before(testcase.d2)
			if got != testcase.before {
				t.Errorf("%v is before %v, want: %v, but got: %v",
					testcase.d1, testcase.d2, testcase.before, got)
			}
		})
	}
}

func TestDate_Before2(t *testing.T) {
	d1 := today()
	time.Sleep(time.Second)
	d2 := today()
	if d1.Before(d2) {
		t.Errorf("%v(%s) is NOT before %v(%s)", d1, d1.Time, d2, d2.Time)
	}
}

func TestDate_After(t *testing.T) {
	for _, testcase := range beforeAfter {
		t.Run("", func(t *testing.T) {
			got := testcase.d1.After(testcase.d2)
			if got != testcase.after {
				t.Errorf("%v is after %v, want: %v, but got: %v",
					testcase.d1, testcase.d2, testcase.after, got)
			}
		})
	}
}

func TestDate_After2(t *testing.T) {
	d1 := today()
	time.Sleep(time.Second)
	d2 := today()
	if d2.After(d1) {
		t.Errorf("%v(%s) is NOT after %v(%s)", d2, d2.Time, d1, d1.Time)
	}
}

func TestTruncate(t *testing.T) {
	now := time.Now()
	trunc := now.Truncate(24 * time.Hour)

	fmt.Println("now:", now)
	fmt.Println("trunc:", trunc)
}

func TestRound(t *testing.T) {
	now := time.Now()
	round := now.Round(24 * time.Hour)

	fmt.Println("now:", now)
	fmt.Println("round:", round)
}

func TestTruncate2(t *testing.T) {
	now := time.Now()
	// Truncate it to the beginning of the day (midnight)
	trunc := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	fmt.Println("now:", now)
	fmt.Println("trunc:", trunc)
}
