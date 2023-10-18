package godate

import "fmt"

func ExampleDate_StringCN() {
	date := MustDate(2023, 10, 18)
	s := date.StringCN()
	fmt.Println(s)

	// Output:
	// 2023年10月18日
}
