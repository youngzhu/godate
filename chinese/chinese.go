package chinese

import (
	"github.com/youngzhu/godate"
	"log"
	"strconv"
)

// 中国特色的调休制度
// TODO 从官方网站查询中国节假日和调休信息

type cnDateSlice []godate.Date

func (c cnDateSlice) test(date godate.Date) bool {
	for _, d := range c {
		if d.IsTheSameDay(date) {
			return true
		}
	}
	return false
}

type chineseDate struct {
	// key: year
	holidays    map[int]cnDateSlice
	extWorkdays map[int]cnDateSlice
}

var cd *chineseDate

func newChineseDate() *chineseDate {
	bc := new(chineseDate)

	bc.holidays = make(map[int]cnDateSlice)
	bc.extWorkdays = make(map[int]cnDateSlice)

	return bc
}

// 不要这个了
// 引用这个包的项目也会执行，如果报错，单元测试都走不下去
//func init() {
//	cd = newChineseDate()
//
//	fetcher = timorFetcher{}
//
//	GetHolidays(godate.Today().Year())
//}

// 校验是否有数据
// 如果没有，则获取并填充
func (c *chineseDate) check(year int) {
	if _, ok := cd.holidays[year]; ok {
		log.Println("cached:", year)
		return
	}

	yearStr := strconv.Itoa(year)
	log.Println("fetch for year:", yearStr)
	holidays, extWorkdays, err := fetchData(yearStr)
	if err != nil {
		log.Fatal(err) // 不会再往下执行了
	}

	log.Println("fetch DONE for year:", yearStr)

	cd.holidays[year] = holidays
	cd.extWorkdays[year] = extWorkdays
}

func GetHolidays(year int) []godate.Date {
	cd.check(year)

	return cd.holidays[year]
}

func GetExtWorkdays(year int) []godate.Date {
	cd.check(year)

	return cd.extWorkdays[year]
}

// IsHoliday 是否中国节假日
func IsHoliday(date godate.Date) bool {
	GetHolidays(date.Year())
	return cd.holidays[date.Year()].test(date)
}

// IsExtWorkday 是否中国式调班日（额外工作日）
func IsExtWorkday(date godate.Date) bool {
	GetExtWorkdays(date.Year())
	return cd.extWorkdays[date.Year()].test(date)
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
