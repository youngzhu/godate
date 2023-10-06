package godate

import (
	"fmt"
	"time"
)

const layoutDate0 = "2006-01-02"

type Date struct {
	time.Time
}

func today() Date {
	now := time.Now()
	// Truncate it to the beginning of the day (midnight)
	d := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return Date{d}

	//return Date{time.Now()}
}

func NewDate() Date {
	return today()
}

// NewDateYMD
// month: 1-12
func NewDateYMD(year, month, day int) (Date, error) {
	s := fmt.Sprintf("%4d-%02d-%02d", year, month, day)
	date, err := time.Parse(layoutDate0, s)
	if err != nil {
		return Date{}, err
	}
	return Date{date}, nil
}

func MustDate(year, month, day int) Date {
	date, _ := NewDateYMD(year, month, day)
	return date
}

func Today() Date {
	return today()
}

func Tomorrow() Date {
	t, _ := today().AddDay(1)
	return t
}

func Yesterday() Date {
	y, _ := today().SubDay(1)
	return y
}

func (d Date) String() string {
	return d.Format(layoutDate0)
}

func (d Date) Month() Month {
	return Month(d.Time.Month())
}

const (
	Day = time.Hour * 24
)

// AddDay returns the date d+days.
func (d Date) AddDay(days int) (Date, error) {
	switch {
	case days < 0:
		return Date{}, fmt.Errorf("days<0, use d.SubDay()")
	case days == 0:
		return d, nil
	default:
		return Date{d.Time.Add(Day * time.Duration(days))}, nil
	}
}

// SubDay returns the date d-days.
func (d Date) SubDay(days int) (Date, error) {
	switch {
	case days < 0:
		return Date{}, fmt.Errorf("days<0, use d.AddDay()")
	case days == 0:
		return d, nil
	default:
		return Date{d.Time.Add(Day * time.Duration(-days))}, nil
	}
}

// MonthDay returns the day of the month specified by d.
func (d Date) MonthDay() int {
	return d.Time.Day()
}

// Weekday returns the day of the week specified by d.
func (d Date) Weekday() Weekday {
	return Weekday(d.Time.Weekday())
}

func (d Date) IsTheSameDay(dd Date) bool {
	return d.Year() == dd.Year() && d.Month() == dd.Month() && d.Day() == dd.Day()
}

func (d Date) Before(x Date) bool {
	return d.Time.Before(x.Time)
}

func (d Date) After(x Date) bool {
	return d.Time.After(x.Time)
}
