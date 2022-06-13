package godate

import (
	"testing"
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

func newDateYMD(year, month, day int) Date {
	d, _ := NewDateYMD(year, month, day)
	return d
}

var addDayTests = []struct {
	date   Date
	days   int
	result Date
}{
	{newDateYMD(2022, 6, 18), 0, newDateYMD(2022, 6, 18)},
	{newDateYMD(2022, 6, 18), 1, newDateYMD(2022, 6, 19)},
	{newDateYMD(2022, 6, 18), 2, newDateYMD(2022, 6, 20)},
}

func TestDate_AddDay(t *testing.T) {
	for _, tt := range addDayTests {
		date, _ := tt.date.AddDay(tt.days)
		if !date.IsTheSameDay(tt.result) {
			t.Errorf("got %s; want %s", date, tt.result)
		}
	}
}
