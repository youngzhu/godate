package fetch

import "github.com/youngzhu/godate/chinese"

var fetcher Fetcher

func init() {
	fetcher = timorFetcher{}
}

type Fetcher interface {
	Fetch(year string) (chinese.Holidays, chinese.ExtWorkdays, error)
}

func Run(year string) (chinese.Holidays, chinese.ExtWorkdays, error) {
	return fetcher.Fetch(year)
}
