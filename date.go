package godate

import (
	"fmt"
	gc "github.com/youngzhu/go2chinese"
	"time"
)

const (
	layoutDate0 = "2006-01-02"

	Day = 24 * time.Hour
)

type Date struct {
	time.Time
	Name          string // 日期描述，如节假日、周末等
	OffdayInChina bool   // 在中国是否为节假日
}

func newDate(d time.Time) Date {
	return Date{
		Time: d,
	}
}

func today() Date {
	now := time.Now()
	// Truncate it to the beginning of the day (midnight)
	d := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return newDate(d)

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
	return newDate(date), nil
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

func (d Date) StringCN() string {
	return gc.Date(d.Time)
}

func (d Date) FullStringCN() string {
	return gc.Date(d.Time) + "，" + gc.Weekday(d.Time)
}

func (d Date) Month() Month {
	return Month(d.Time.Month())
}

// AddDay returns the date d+days.
func (d Date) AddDay(days int) (Date, error) {
	switch {
	case days < 0:
		return Date{}, fmt.Errorf("days<0, use d.SubDay()")
	case days == 0:
		return d, nil
	default:
		return newDate(d.Time.Add(Day * time.Duration(days))), nil
	}
}

func (d Date) AddDaySafe(days int) Date {
	addedDate, _ := d.AddDay(days)
	return addedDate
}

// SubDay returns the date d-days.
func (d Date) SubDay(days int) (Date, error) {
	switch {
	case days < 0:
		return Date{}, fmt.Errorf("days<0, use d.AddDay()")
	case days == 0:
		return d, nil
	default:
		return newDate(d.Time.Add(Day * time.Duration(-days))), nil
	}
}

// MonthDay returns the day of the month specified by d.
func (d Date) MonthDay() int {
	return d.Time.Day()
}

// Weekday returns the day of the week specified by d.
func (d Date) Weekday() Weekday {
	weekday := int(d.Time.Weekday())
	if weekday == 0 {
		return Sunday
	} else {
		return Weekday(weekday)
	}
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

/*
形成引用环了。这个判断逻辑都放到chinese包里吧

// IsOffDayInChina 在中国是否放假日
// 放假应满足（满足其一即可）：
// 1. 周末 且 没有调班
// 2. 中国节假日
//
// 上班应满足（满足其一即可）：
// 1. 工作日（周一至周五） 且 不是中国节假日
// 2. 调班
func (d Date) IsOffDayInChina() bool {
	if chinese.IsHoliday(d) {
		return true
	}

	return d.IsWeekend() && !chinese.IsExtWorkday(d)
}

// IsWorkDayInChina 在中国是否上班日
func (d Date) IsWorkDayInChina() bool {

	return !d.IsOffDayInChina()
}
*/
