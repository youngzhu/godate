package godate

import "fmt"

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var longWeekdayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortWeekdayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Weekday) String() string {
	if Sunday <= d && d <= Saturday {
		return longWeekdayNames[d]
	}
	return fmt.Sprintf("!Weekday(\"%d\")", d)
}

func (d Date) Workdays() (workdays []Date) {
	w := d.Weekday()
	switch w {
	case Sunday:
		for i := Monday; i <= Friday; i++ {
			day, _ := d.AddDay(int(i))
			workdays = append(workdays, day)
		}
	case Saturday:
		for i := Friday; i >= Monday; i-- {
			day, _ := d.SubDay(int(i))
			workdays = append(workdays, day)
		}
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
