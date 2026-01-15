package chinese

import (
	"github.com/youngzhu/godate"
	"log"
	"testing"
)

func TestIsOffDayInChina(t *testing.T) {
	testCases := []struct {
		date godate.Date
		want bool
	}{
		{godate.MustDate(2024, 1, 1), true},
		{godate.MustDate(2024, 1, 24), false},
		{godate.MustDate(2024, 2, 9), false},
		{godate.MustDate(2024, 2, 10), true},
		{godate.MustDate(2024, 2, 11), true},
		{godate.MustDate(2024, 2, 12), true},
		{godate.MustDate(2024, 2, 13), true},
		{godate.MustDate(2024, 2, 14), true},
		{godate.MustDate(2024, 2, 15), true},
		{godate.MustDate(2024, 2, 16), true},
		{godate.MustDate(2024, 2, 17), true},
		{godate.MustDate(2024, 2, 18), false},
		{godate.MustDate(2025, 1, 1), true},
		{godate.MustDate(2025, 10, 1), false},
	}

	for _, testcase := range testCases {
		t.Run("", func(t *testing.T) {
			got := IsOffDayInChina(testcase.date)
			if got != testcase.want {
				t.Errorf("%v is off day in China? want: %v, but got: %v",
					testcase.date, testcase.want, got)
			}
		})
	}
}

func TestIsWorkDayInChina(t *testing.T) {
	testCases := []struct {
		date godate.Date
		want bool
	}{
		{godate.MustDate(2024, 1, 1), false},
		{godate.MustDate(2024, 2, 4), true},
		{godate.MustDate(2024, 2, 9), true},
		{godate.MustDate(2024, 2, 17), false},
		{godate.MustDate(2024, 2, 18), true},
		{godate.MustDate(2026, 1, 1), false},
		{godate.MustDate(2026, 1, 3), false},
		{godate.MustDate(2026, 1, 4), true},
	}

	for _, testcase := range testCases {
		t.Run("", func(t *testing.T) {
			got := IsWorkDayInChina(testcase.date)
			if got != testcase.want {
				t.Errorf("%v is work day in China? want: %v, but got: %v",
					testcase.date, testcase.want, got)
			}
		})
	}
}

func newDate(year, month, day int, name string) CNDate {
	date := godate.MustDate(year, month, day)

	return CNDate{
		Date: date,
		Name: name,
	}
}

func TestGetOffdaysByRange(t *testing.T) {
	offdays, _ := GetOffdaysByRange("20260101", "20260111")
	expectedOffdays := []CNDate{
		newDate(2026, 1, 1, "元旦"),
		newDate(2026, 1, 2, "元旦"),
		newDate(2026, 1, 3, "元旦"),
		newDate(2026, 1, 10, "周末"),
	}

	if len(offdays) != len(expectedOffdays) {
		t.Fatalf("expected %d offdays, got %d", len(expectedOffdays), len(offdays))
	}

	for i, d := range expectedOffdays {
		if !offdays[i].IsTheSameDay(d.Date) || offdays[i].Name != d.Name {
			t.Errorf("expected offday %v(%s) at index %d, got %v(%s)", d, d.Name, i, offdays[i], offdays[i].Name)
		}
	}
}

func TestGetOffdaysOfYear_count(t *testing.T) {
	offdays, _ := GetOffdaysOfYear(2026)

	// 打印一下
	for i, offday := range offdays {
		log.Println(i, offday, offday.Name)
	}

	expectedOffdays := 117
	if len(offdays) != expectedOffdays {
		t.Fatalf("expected %d offdays, got %d", expectedOffdays, len(offdays))
	}

}

func TestCNDate_String(t *testing.T) {
	today := NewCNDate(godate.Today())
	log.Println(today.String())
}

func TestCNDate_AddDay(t *testing.T) {
	day1231 := NewCNDate(godate.MustDate(2025, 12, 31))
	day0101, _ := day1231.AddDay(1)

	wantName := "元旦"
	wantOffday := true
	if day0101.Name != wantName {
		t.Errorf("expected name %s, got %s", wantName, day0101.Name)
	}
	if day0101.Offday != wantOffday {
		t.Errorf("expected offday %v, got %v", wantOffday, day0101.Offday)
	}

	//today := NewCNDate(godate.Today())
	//tomorrow, _ := today.AddDay(1)
	//
	//log.Printf("today: %v, tomorrow: %v", today, tomorrow)
}

func TestCNDate_NextDay(t *testing.T) {
	day1231 := NewCNDate(godate.MustDate(2025, 12, 31))
	day0101 := day1231.NextDay()

	wantName := "元旦"
	wantOffday := true
	if day0101.Name != wantName {
		t.Errorf("expected name %s, got %s", wantName, day0101.Name)
	}
	if day0101.Offday != wantOffday {
		t.Errorf("expected offday %v, got %v", wantOffday, day0101.Offday)
	}
}

// func TestTimorFetcher_Fetch_fail(t *testing.T) {
// 	GetHolidays(2023)
// }
