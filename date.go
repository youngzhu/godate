package godate

import (
	"fmt"
	"time"
)

const layoutDate0 = "2006-01-02"

type Date struct {
	time.Time
}

func NewDate() Date {
	return Date{time.Now()}
}

func NewDateYMD(year, month, day int) (Date, error) {
	s := fmt.Sprintf("%4d-%02d-%02d", year, month, day)
	date, err := time.Parse(layoutDate0, s)
	if err != nil {
		return Date{}, err
	}
	return Date{date}, nil
}

func (d Date) String() string {
	return d.Format(layoutDate0)
}
