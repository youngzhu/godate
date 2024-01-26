package chinese

var fetcher Fetcher

type Fetcher interface {
	Fetch(year string) (cnDateSlice, cnDateSlice, error)
}

func fetchData(year string) (cnDateSlice, cnDateSlice, error) {
	return fetcher.Fetch(year)
}
