package chinese

import (
	"github.com/youngzhu/godate"
	"log"
	"strconv"
)

// 中国特色的调休制度
// TODO 从官方网站查询中国节假日和调休信息

type (
	Holidays    []godate.Date // 中国节假日
	ExtWorkdays []godate.Date // 中国式调休日
)

type chineseDate struct {
	// key: year
	holidays    map[int]Holidays
	extWorkdays map[int]ExtWorkdays
}

var cd *chineseDate

func newChineseDate() *chineseDate {
	bc := new(chineseDate)

	bc.holidays = make(map[int]Holidays)
	bc.extWorkdays = make(map[int]ExtWorkdays)

	return bc
}

func init() {
	cd = newChineseDate()
}

func GetHolidays(year int) []godate.Date {
	if holidays, ok := cd.holidays[year]; ok {
		log.Println("cached:", year)
		return holidays
	}

	yearStr := strconv.Itoa(year)
	log.Println("search for year", yearStr)
	holidays, extWorkdays, err := fetchData(yearStr)
	if err != nil {
		return nil
	}

	cd.holidays[year] = holidays
	cd.extWorkdays[year] = extWorkdays

	return cd.holidays[year]
}

func GetExtWorkdays(year int) []godate.Date {
	if extWorkdays, ok := cd.extWorkdays[year]; ok {
		log.Println("cached data:", year)
		return extWorkdays
	}

	yearStr := strconv.Itoa(year)
	log.Println("search for year", yearStr)
	holidays, extWorkdays, err := fetchData(yearStr)
	if err != nil {
		return nil
	}

	cd.holidays[year] = holidays
	cd.extWorkdays[year] = extWorkdays

	return cd.extWorkdays[year]
}

// IsHoliday 是否中国节假日
func IsHoliday(date godate.Date) bool {
	for _, d := range GetHolidays(date.Year()) {
		if d.IsTheSameDay(date) {
			return true
		}
	}

	return false
}

// IsExtWorkday 是否中国式调班日
func IsExtWorkday(date godate.Date) bool {
	for _, d := range GetExtWorkdays(date.Year()) {
		if d.IsTheSameDay(date) {
			return true
		}
	}

	return false
}

// IsOffDayInChina 在中国是否放假日
// 放假应满足（满足其一即可）：
// 1. 周末 且 没有调班
// 2. 中国节假日
//
// 上班应满足（满足其一即可）：
// 1. 工作日（周一至周五） 且 不是中国节假日
// 2. 调班
func IsOffDayInChina(d godate.Date) bool {
	if IsHoliday(d) {
		return true
	}

	return d.IsWeekend() && !IsExtWorkday(d)
}

// IsWorkDayInChina 在中国是否上班日
func IsWorkDayInChina(d godate.Date) bool {

	return !IsOffDayInChina(d)
}
