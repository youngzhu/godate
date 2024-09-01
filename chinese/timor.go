package chinese

import (
	"encoding/json"
	"fmt"
	"github.com/youngzhu/godate"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// 通过 https://timor.tech/api/holiday/{year} 获取数据

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

func (f timorFetcher) Fetch(year string) (cnDateSlice, cnDateSlice, error) {
	//if val, ok := timorData[year]; ok {
	//	log.Println("预处理数据")
	//	return readFromJson(strings.NewReader(val))
	//}
	var holidays, workdays cnDateSlice

	// 先从本地获取
	//holidays, workdays, err := readFromLocal(year)
	//if err == nil {
	//	return holidays, workdays, nil
	//}

	// 通过API获取

	url := fmt.Sprintf("https://cdn.jsdelivr.net/gh/youngzhu/haohanbang/wodc/%s.json", year)
	resp, err := http.Get(url)
	if err == nil {
		holidays, workdays, err = readFromJson(resp.Body)
		if err == nil {
			return holidays, workdays, nil
		}
	}

	return nil, nil, fmt.Errorf("暂未实现")
}

func readFromJson(jsonContent io.Reader) (cnDateSlice, cnDateSlice, error) {
	// Decode the file into a slice of pointers
	// to timorResult values.
	var result *timorResult
	err := json.NewDecoder(jsonContent).Decode(&result)
	if err != nil {
		return nil, nil, err
	}

	holidays := make(cnDateSlice, 0, len(result.Holidays))
	extWorkdays := make(cnDateSlice, 0, len(result.Holidays))

	for _, day := range result.Holidays {
		date, err := godate.Parse(day.Date)
		if err != nil {
			return nil, nil, err
		}
		if day.Holiday {
			holidays = append(holidays, date)
		} else {
			extWorkdays = append(extWorkdays, date)
		}
	}

	// We don't need to check for errors, the caller can do this.
	return holidays, extWorkdays, err
}

const rootPath = "data/timor"

func readFromLocal(year string) (cnDateSlice, cnDateSlice, error) {
	// Open the file.
	file, err := os.Open(filepath.Join(rootPath, year+".json"))
	if err != nil {
		return nil, nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	return readFromJson(file)
}
