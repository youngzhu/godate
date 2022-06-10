package godate

func (d Date) IsLeapYear() bool {
	year := d.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
