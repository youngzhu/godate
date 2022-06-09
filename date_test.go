package godate

import "testing"

func TestNewDate(t *testing.T) {
	t.Log(NewDate())
}

func TestNewDateYMD(t *testing.T) {
	t.Log(NewDateYMD(2000, 1, 1))
}
