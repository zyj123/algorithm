package top100InBigFile

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"os"
	"strings"
)

func generateOneGFile() {
	f, err := os.Create("words.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(f)

	words := readWords()
	for i := 0; i < 1<<27; i++ {
		rand.Seed(int64(i))
		word := words[rand.Intn(len(words))]
		_, err = w.WriteString(word + "\n")
		if err != nil {
			panic(err)
		}
		fmt.Printf("write word %s success...\n", word)
	}
}

func readWords() []string {
	f, err := excelize.OpenFile("./5000.xlsx")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()
	rows, err := f.GetRows("工作表1")
	if err != nil {
		panic(err)
	}
	words := make([]string, 0, 5000)
	for _, row := range rows {
		words = append(words, strings.TrimSpace(row[1]))
	}

	return words
}
