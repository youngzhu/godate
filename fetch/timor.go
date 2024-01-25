package fetch

import (
	"encoding/json"
	"github.com/youngzhu/godate"
	"github.com/youngzhu/godate/chinese"
	"os"
	"path/filepath"
)

// 通过 https://timor.tech/api/holiday/ 获取数据

type timorResult struct {
	Code     int `json:"code"`
	Holidays []struct {
		Holiday bool   `json:"holiday"`
		Name    string `json:"name"`
		Wage    int    `json:"wage"`
		Date    string `json:"date"`
		Rest    int    `json:"rest,omitempty"`
		After   bool   `json:"after,omitempty"`
		Target  string `json:"target,omitempty"`
	} `json:"holidays"`
}

type timorFetcher struct{}

func (f timorFetcher) Fetch(year string) (chinese.Holidays, chinese.Workdays, error) {
	// 先从本地获取
	holidays, workdays, err := readFromLocal(year)
	if err == nil {
		return holidays, workdays, nil
	}

	// 通过API获取

	return nil, nil, nil
}

const rootPath = "data/timor"

func readFromLocal(year string) (chinese.Holidays, chinese.Workdays, error) {
	// Open the file.
	file, err := os.Open(filepath.Join(rootPath, year+".json"))
	if err != nil {
		return nil, nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to timorResult values.
	var result *timorResult
	err = json.NewDecoder(file).Decode(&result)
	if err != nil {
		return nil, nil, err
	}

	holidays := make(chinese.Holidays, 0, len(result.Holidays))
	workdays := make(chinese.Workdays, 0, len(result.Holidays))

	for _, day := range result.Holidays {
		date, err := godate.Parse(day.Date)
		if err != nil {
			return nil, nil, err
		}
		if day.Holiday {
			holidays = append(holidays, date)
		} else {
			workdays = append(workdays, date)
		}
	}

	// We don't need to check for errors, the caller can do this.
	return holidays, workdays, err
}
