package godate

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (d Date) IsLeapYear() bool {
	year := d.Year()
	return isLeap(year)
}

func (d Date) YearDays() int {
	if d.IsLeapYear() {
		return int(daysBefore[December]) + 1
	}
	return int(daysBefore[December])
}
