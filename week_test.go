package godate

import "testing"

func TestWeekday_String(t *testing.T) {
	t.Log(Monday.String())
	t.Log(Weekday(14).String())
}
