package chinese

import (
	"fmt"
	"testing"
)

func TestTimor_readFromLocal(t *testing.T) {
	holidays, workdays, err := readFromLocal("2024")

	if err != nil {
		t.Error(err)
	}

	if len(holidays) == 0 {
		t.Error("holidays should not empty.")
	}

	if len(workdays) == 0 {
		t.Error("workdays should not empty.")
	}

	fmt.Println("holidays:", holidays)
	fmt.Println("workdays:", workdays)
}

func TestTimorFetcher_Fetch(t *testing.T) {
	_timorFetcher := timorFetcher{}
	holidays, workdays, err := _timorFetcher.Fetch("2023")

	if err != nil {
		t.Error(err)
	}

	if len(holidays) == 0 {
		t.Error("holidays should not empty.")
	}

	if len(workdays) == 0 {
		t.Error("workdays should not empty.")
	}

	fmt.Println("holidays:", holidays)
	fmt.Println("workdays:", workdays)
}
