package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

}

func modify(nums []int) {
	nums = append(nums, 1)
	nums[0] = 999
}

func main2() {

	boy := []string{
		"boy_weight.csv",
		"boy_height.csv",
		"boy_head.csv",
	}

	var data []Source
	for _, f := range boy {
		data = append(data, readCSV(f)...)
	}

	d, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d) + ",")
}

type Source struct {
	Source [][]any `json:"source"`
}

func readCSV(file string) []Source {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var (
		p3  [][]any
		p10 [][]any
		p25 [][]any
		p50 [][]any
		p75 [][]any
		p90 [][]any
		p97 [][]any
	)

	for i, record := range records {
		if i%5 != 0 && i%24 != 0 {
			continue
		}
		p3 = append(p3, []any{i, record[1]})
		p10 = append(p10, []any{i, record[2]})
		p25 = append(p25, []any{i, record[3]})
		p50 = append(p50, []any{i, record[4]})
		p75 = append(p75, []any{i, record[5]})
		p90 = append(p90, []any{i, record[6]})
		p97 = append(p97, []any{i, record[7]})
	}

	return []Source{
		Source{Source: p3},
		Source{Source: p10},
		Source{Source: p25},
		Source{Source: p50},
		Source{Source: p75},
		Source{Source: p90},
		Source{Source: p97},
	}
}
