package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)
	
var f = flag.String("f", "0", "which columns to print")
var d = flag.String("d", "\t", "delimiter for splitting lines")
var s = flag.Bool("s", false, "show only lines with delimiter")

func parseColumns(query string) (indexes []int, err error) {
	numberStrings := strings.Split(strings.ReplaceAll(query, ", ", ","), ",")
	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return indexes, err
		}
		indexes = append(indexes, number)
	}
	return indexes, err
}

func Cut(data string, f string, d string, s bool) (result string, err error) {

	cols, err := parseColumns(f)
	if err != nil {
		return "", err
	}

	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")
	for _, line := range lines {
		if strings.Contains(line, d) {
			lineSlice := strings.Split(line, d)
			for i, index := range cols {
				if index >= len(lineSlice) {
					continue
				}
				result += lineSlice[index]

				if i != len(cols)-1 { // To avoid trailing space
					result += " "
				}
			}
			result += "\n"
		} else if !s {
			result += line + "\n"
		}
	}
	return result, nil
}

func main() {
	// Init
	flag.Parse()
	args := flag.Args()
	src := args[0]
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("No such file")
		os.Exit(1)
	}

	// Call cut
	if result, err := Cut(string(data), *f, *d, *s); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}