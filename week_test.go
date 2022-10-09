package godate

import (
	"reflect"
	"testing"
)

func TestWeekday_String(t *testing.T) {
	t.Log(Monday.String())
	t.Log(Weekday(14).String())
}

var (
	oneWeek = []Date{
		newDateYMD(2022, 6, 12),
		newDateYMD(2022, 6, 13),
		newDateYMD(2022, 6, 14),
		newDateYMD(2022, 6, 15),
		newDateYMD(2022, 6, 16),
		newDateYMD(2022, 6, 17),
		newDateYMD(2022, 6, 18),
	}
	workdays = oneWeek[Monday:Saturday]
)

func TestDate_Workdays(t *testing.T) {
	for i, date := range oneWeek {
		got := date.Workdays()
		if !reflect.DeepEqual(got, workdays) {
			t.Errorf("%s, got: %v, want: %v", Weekday(i), got, workdays)
		}
	}
}

var nextWorkdays = []struct {
	date        Date
	nextWorkday Date
}{
	{newDateYMD(2022, 10, 9), newDateYMD(2022, 10, 10)},
	{newDateYMD(2022, 10, 10), newDateYMD(2022, 10, 11)},
	{newDateYMD(2022, 10, 11), newDateYMD(2022, 10, 12)},
	{newDateYMD(2022, 10, 13), newDateYMD(2022, 10, 14)},
	{newDateYMD(2022, 10, 14), newDateYMD(2022, 10, 17)},
	{newDateYMD(2022, 10, 15), newDateYMD(2022, 10, 17)},
	{newDateYMD(2022, 10, 31), newDateYMD(2022, 11, 1)},
}

func TestDate_NextWorkday(t *testing.T) {
	for _, tt := range nextWorkdays {
		next := tt.date.NextWorkday()
		if next != tt.nextWorkday {
			t.Errorf("got %s; want %s", next, tt.nextWorkday)
		}
	}
}
