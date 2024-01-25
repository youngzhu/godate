package chinese

import (
	"github.com/youngzhu/godate"
	"log"
)

// 中国特色的调休制度
// TODO 从官方网站查询中国节假日和调休信息

type (
	Holidays []godate.Date // 中国节假日
	Workdays []godate.Date // 中国工作日
)

type chineseDate struct {
	// key: year
	holidays map[int][]godate.Date
}

var cd *chineseDate

func newChineseDate() *chineseDate {
	bc := new(chineseDate)

	bc.holidays = make(map[int][]godate.Date)

	return bc
}

func init() {
	cd = newChineseDate()
}

func ChineseHolidays(year int) []godate.Date {
	if holidays, ok := cd.holidays[year]; ok {
		log.Println("cached:", year)
		return holidays
	}
	log.Println("search for year", year)

	return nil
}
