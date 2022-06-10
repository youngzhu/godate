package godate

import "testing"

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
