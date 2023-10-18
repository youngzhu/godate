package godate

import "fmt"

func ExampleDate_StringCN() {
	date := MustDate(2023, 10, 18)
	s := date.StringCN()
	fmt.Println(s)

	// Output:
	// 2023年10月18日
}

func ExampleDate_FullStringCN() {
	date := MustDate(2023, 10, 18)
	s := date.FullStringCN()
	fmt.Println(s)

	// Output:
	// 2023年10月18日，星期三
}
