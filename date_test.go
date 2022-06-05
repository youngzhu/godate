package godate

import (
	"testing"
	"time"
)

var (
	d1 Date1
	d2 Date2
	d3 Date3
	d4 Date4
	d5 Date5
	d6 Date6
)

func TestDate(t *testing.T) {
	t.Logf("d1: %+v", d1)
	t.Logf("d1: %+v", Date1{})

	t.Logf("d2: %+v", d2)
	t.Logf("d2: %+v", Date2{})

	t.Logf("d3: %+v", d3)
	t.Logf("d3: %+v", Date3{})

	t.Logf("d4: %+v", d4)
	t.Logf("d4: %+v", Date4{})

	t.Logf("d5: %+v", d5)
	t.Logf("d5: %+v", Date5{})

	t.Logf("d6: %+v", d6)
	//t.Logf("d6: %+v", Date6{})
}

func TestMethod(t *testing.T) {
	d1 = Date1{time.Now()}
	d11 := d1.AddDate(0, 0, 1)
	d12 := d1.Add(time.Hour)
	t.Logf("d1: %#v", d1)
	t.Logf("d11: %#v", d11)
	t.Logf("d12: %#v", d12)
	t.Logf("d12: %#v", d1.GoString())
	t.Logf("d12: %#v", d1.YearDay())
	t.Logf("d12: %#v", d1.Weekday())

	//d2.datetime

	//d3=Date3{time.Now()}
	//d33 := d3.AddDate(0, 0, 3)
	//t.Logf("d3: %+v", d3)
	//t.Logf("d33: %+v", d33)

	//d55 := d5.AddDate(0, 0, 5)
	//t.Logf("d5: %+v", d5)
	//t.Logf("d55: %+v", d55)

	//d66 := d6.AddDate(0, 0, 6)
	//t.Logf("d6: %+v", d6)
	//t.Logf("d66: %+v", d66)
}
