package godate

import "fmt"

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var shortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string {
	if January <= m && m <= December {
		return longMonthNames[m-1]
	}
	return fmt.Sprintf("!Month(\"%d\")", m)
}

func (m Month) LongName() string {
	if January <= m && m <= December {
		return longMonthNames[m-1]
	}
	return fmt.Sprintf("!Month(\"%d\")", m)
}

func (m Month) ShortName() string {
	if January <= m && m <= December {
		return shortMonthNames[m-1]
	}
	return fmt.Sprintf("!Month(\"%d\")", m)
}

func (m Month) IntValue() int {
	return int(m)
}
