package chinese_test

import (
	"fmt"
	"github.com/youngzhu/godate"
	"github.com/youngzhu/godate/chinese"
)

var (
	jan1 = godate.MustDate(2024, 1, 1) // 元旦
	feb4 = godate.MustDate(2024, 2, 4) // 2024年第一个调班
)

func ExampleIsOffDayInChina() {
	fmt.Println(chinese.IsOffDayInChina(jan1))
	fmt.Println(chinese.IsOffDayInChina(feb4))

	// Output:
	// true
	// false
}

func ExampleIsWorkDayInChina() {
	fmt.Println(chinese.IsWorkDayInChina(jan1))
	fmt.Println(chinese.IsWorkDayInChina(feb4))

	// Output:
	// false
	// true
}

func ExampleGetHolidays() {
	for _, d := range chinese.GetHolidays(2024) {
		fmt.Println(d)
	}

	// Output:
	//2024-01-01
	//2024-02-10
	//2024-02-11
	//2024-02-12
	//2024-02-13
	//2024-02-14
	//2024-02-15
	//2024-02-16
	//2024-02-17
	//2024-04-04
	//2024-04-05
	//2024-04-06
	//2024-05-01
	//2024-05-02
	//2024-05-03
	//2024-05-04
	//2024-05-05
	//2024-06-08
	//2024-06-09
	//2024-06-10
	//2024-09-15
	//2024-09-16
	//2024-09-17
	//2024-10-01
	//2024-10-02
	//2024-10-03
	//2024-10-04
	//2024-10-05
	//2024-10-06
	//2024-10-07
}

func ExampleGetExtWorkdays() {
	for _, d := range chinese.GetExtWorkdays(2024) {
		fmt.Println(d)
	}

	// Output:
	//2024-02-04
	//2024-02-18
	//2024-04-07
	//2024-04-28
	//2024-05-11
	//2024-09-14
	//2024-09-29
	//2024-10-12
}

func ExampleIsHoliday() {
	fmt.Println(chinese.IsHoliday(jan1)) // 元旦

	sat := godate.MustDate(2024, 1, 27) // 周末放假，但不是中国节假日
	fmt.Println(chinese.IsHoliday(sat))

	// Output:
	// true
	// false
}

func ExampleIsExtWorkday() {
	fri := godate.MustDate(2024, 1, 26) // 周五，正常工作日
	fmt.Println(chinese.IsExtWorkday(fri))

	// 调班
	fmt.Println(chinese.IsExtWorkday(feb4))

	// Output:
	// false
	// true
}
