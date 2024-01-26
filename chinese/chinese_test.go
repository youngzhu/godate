package chinese

import (
	"github.com/youngzhu/godate"
	"testing"
)

func TestIsOffDayInChina(t *testing.T) {
	testCases := []struct {
		date godate.Date
		want bool
	}{
		{godate.MustDate(2024, 1, 1), true},
		{godate.MustDate(2024, 1, 24), false},
	}

	for _, testcase := range testCases {
		t.Run("", func(t *testing.T) {
			got := IsOffDayInChina(testcase.date)
			if got != testcase.want {
				t.Errorf("%v is Chinese holiday? want: %v, but got: %v",
					testcase.date, testcase.want, got)
			}
		})
	}
}
