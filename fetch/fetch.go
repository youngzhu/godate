package fetch

import "github.com/youngzhu/godate/chinese"

type Fetcher interface {
	Fetch(year int) (chinese.Holidays, chinese.Workdays, error)
}
