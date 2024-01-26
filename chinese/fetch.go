package chinese

var fetcher Fetcher

func init() {
	fetcher = timorFetcher{}
}

type Fetcher interface {
	Fetch(year string) (Holidays, ExtWorkdays, error)
}

func fetchData(year string) (Holidays, ExtWorkdays, error) {
	return fetcher.Fetch(year)
}
