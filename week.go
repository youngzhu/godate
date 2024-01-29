package godate

import "fmt"

// A Weekday representing the 7 days of the week.
// In addition to the textual name, each day-of-week has an int value.
// The int value follows the ISO-8601 standard, from 1 (Monday) to 7 (Sunday).
type Weekday int

const (
	Monday Weekday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var longWeekdayNames = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

var shortWeekdayNames = []string{
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
	"Sun",
}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (w Weekday) String() string {
	if Monday <= w && w <= Sunday {
		return longWeekdayNames[w]
	}
	return fmt.Sprintf("!Weekday(\"%d\")", w)
}

// IntValue returns the day-of-week to represent, from 1 (Monday) to 7 (Sunday)
func (w Weekday) IntValue() int {
	return int(w)
}

func (d Date) Workdays() (workdays []Date) {
	w := d.Weekday()
	switch w {
	case Saturday, Sunday:
		for i, cnt := int(w)-1, 5; i > 0 && cnt > 0; i, cnt = i-1, cnt-1 {
			day, _ := d.SubDay(int(i))
			workdays = append(workdays, day)
		}
	//case Sunday:
	//	for i := Monday; i <= Friday; i++ {
	//		day, _ := d.AddDay(int(i))
	//		workdays = append(workdays, day)
	//	}
	default:
		workdays = make([]Date, 5)

		wi := int(w) - 1
		workdays[wi] = d

		before, after := int(w-Monday), int(Friday-w)
		for before > 0 {
			day, _ := d.SubDay(before)
			workdays[wi-before] = day
			before--
		}

		for after > 0 {
			day, _ := d.AddDay(after)
			workdays[wi+after] = day
			after--
		}
	}
	return
}

func (d Date) NextWorkday() Date {
	dayOfWeek := d.Weekday()
	var dayToAdd int
	switch dayOfWeek {
	case Friday:
		dayToAdd = 3
	case Saturday:
		dayToAdd = 2
	default:
		dayToAdd = 1
	}

	next, _ := d.AddDay(dayToAdd)
	return next
}

func (d Date) IsWorkday() bool {
	return !d.IsWeekend()
}

func (d Date) IsWeekend() bool {
	weekday := d.Weekday()
	return Saturday == weekday || Sunday == weekday
}
