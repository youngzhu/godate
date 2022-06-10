package godate

import "testing"

func TestMonth(t *testing.T) {
	t.Log(January)
}

func TestMonth_String(t *testing.T) {
	t.Log(January.String())
	t.Log(Month(13).String())
}
