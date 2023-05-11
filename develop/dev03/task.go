package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	nFlag = flag.Bool("n", false, "Нумерованная сортировка")
	rFlag = flag.Bool("r", false, "Обратная сортировка")
	uFlag = flag.Bool("u", false, "Cортировка без повторяющихся строк")
	kFlag = flag.Int("k", 0, "Cортировка с указанием колонки")
)

type dataFile struct {
	Strings    []string
	Words      [][]string
	Space      rune
	MaxColumns int
}

func readFile(data *dataFile, k int, unique bool) error {
	i := 0
	if len(os.Args) < 2 {
		err := errors.New("Не указан файл")
		return err
	}
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	set := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if unique {
			if _, ok := set[line]; !ok {
				set[line] = true
			} else {
				continue
			}
		}
		if k == 0 {
			data.Strings = append(data.Strings, line)
		} else {
			lineWords := strings.Fields(line)
			if data.MaxColumns < len(lineWords) {
				data.MaxColumns = len(lineWords)
			}
			data.Words = append(data.Words, lineWords)
			data.Words[i] = append(data.Words[i], "\n")
			i++
		}
		if err = scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

func simpleSort(data *dataFile, n bool) {
	if n {
		sort.Slice(data.Strings, func(i, j int) bool {
			valueI, _ := strconv.Atoi(data.Strings[i])
			valueJ, _ := strconv.Atoi(data.Strings[j])
			return valueI < valueJ
		})
	} else {
		sort.Strings(data.Strings)
	}
}

func printText(data *dataFile) {
	if data.Words != nil {
		for _, row := range data.Words {
			for _, value := range row {
				switch value {
				case "\n":
					data.Space = 0
				default:
					data.Space = ' '
				}
				fmt.Printf("%s%c", value, data.Space)
			}
		}
	} else {
		for _, value := range data.Strings {
			fmt.Println(value)
		}
	}
}

func sortWithColumn(data *dataFile, k int) {
	if k == 0 || k > data.MaxColumns {
		data.Words = nil
		return
	}
	compare := func(i, j int) bool {
		if data.Words[i][0] == "\n" {
			return true
		} else if data.Words[j][0] == "\n" {
			return false
		}
		return data.Words[i][k-1] < data.Words[j][k-1]
	}
	sort.Slice(data.Words, compare)
}

func reverseSort(data *dataFile, r bool) {
	if r {
		if data.Words != nil {
			sort.Slice(data.Words, func(i, j int) bool {
				return data.Words[i][0] > data.Words[j][0]
			})
		} else {
			sort.Slice(data.Strings, func(i, j int) bool {
				return data.Strings[i] < data.Strings[j]
			})
		}
	}
}

func main() {
	data := dataFile{}
	flag.Parse()
	if err := readFile(&data, *kFlag, *uFlag); err != nil {
		log.Fatal(err)
	}

	simpleSort(&data, *nFlag)
	sortWithColumn(&data, *kFlag)
	reverseSort(&data, *rFlag)
	printText(&data)
}
