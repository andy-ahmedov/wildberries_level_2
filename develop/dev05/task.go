package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var A = flag.Int("A", 0, "print N lines after match")
var B = flag.Int("B", 0, "print N lines before match")
var C = flag.Int("C", 0, "print +- N around match")
var c = flag.Bool("c", false, "print number of matches")
var i = flag.Bool("i", false, "ignore case")
var v = flag.Bool("v", false, "invert search")
var F = flag.Bool("F", false, "exact non-pattern match")
var n = flag.Bool("n", false, "print line index")

func grep(data string, pattern string, A, B, C int, c, i, v, F, n bool) (string, error) {
	isInResult := make(map[int]bool)
	counter := 0

	// Разделяем файл на строки из текста
	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")

	// Преобразование pattern в нижний регистр, если поиск нечувствителен к регистру
	if i {
		pattern = strings.ToLower(pattern)
	}

	for j, line := range lines {

		var (
			matched bool
			err     error
		)

		// Преобразование строки в нижний регистр, если поиск не чувствителен к регистру
		if i {
			line = strings.ToLower(line)
		}

		if !F { // Если искать по pattern
			// Проверка, содержит ли строка pattern
			matched, err = regexp.MatchString(pattern, line)
			if err != nil {
				return "", err
			}
		} else { // Если фиксированный поиск
			matched = strings.Contains(line, pattern)
		}

		if matched != v { // С учетом инвертирования
			counter++
			if A > 0 {
				C = 0                                         
				for k := j; k < min(len(lines), j+A+1); k++ { // min(len(lines), j+A+1) чтобы изжежать выхода за пределы массива
					isInResult[k] = true
				}
			}

			if B > 0 {
				C = 0                               
				for k := max(j-B, 0); k <= j; k++ { // min(j-B, 0) чтобы изжежать выхода за пределы массива	
					isInResult[k] = true
				}
			}

			if C > 0 {
				for k := max(j-C, 0); k < min(len(lines), j+C+1); k++ {
					isInResult[k] = true
				}
			}

			if A == 0 && B == 0 && C == 0 {
				isInResult[j] = true
			}
		}
	}

	if counter == 0 { // If no match found
		return "no match", nil
	}

	result := ""

	if c { 
		result += fmt.Sprintf("found matches: %d", counter) + "\n"
	}

	for index, line := range lines { 
		if isInResult[index] {
			newLine := ""
			if n { 
				newLine += strconv.Itoa(index) + " "
			}
			newLine += line
			result += newLine + "\n"
		}
	}

	return result, nil
}

func main() {
	// Init
	flag.Parse()
	args := flag.Args()
	src := args[0]
	pattern := args[1]
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("No such file")
		os.Exit(1)
	}

	// Вызов grep
	if result, err := grep(string(data), pattern, *A, *B, *C, *c, *i, *v, *F, *n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}