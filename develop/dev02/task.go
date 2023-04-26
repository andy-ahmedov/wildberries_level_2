package main

import (
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся
символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


type edelveyse struct {
	
}

func find_num(str string, index int) int {
	num := 0
	flag := true
	// fmt.Println(strconv.Atoi(string(str[index])))
	for i := index; i < len(str); i++ {
		digit, err := strconv.Atoi(string(str[i]))
		// fmt.Println("Digit = ", digit)
		if err == nil && flag == true {
			num = num*10 + digit
			// fmt.Println("Num = ", num)
		} else {
			flag = false
		}
		if flag == false {
			return num - 1
		}
	}
	//fmt.Println("ETO ", num)
	return num - 1
}

func unpacking_str(str string) string {
	isDigit := false
	new_str := []rune("")
	arr := []rune(str)
	for i, value := range arr {
		_, err := strconv.Atoi(string(value))
		if err == nil && i != 0 {
			if isDigit {
				continue
			}
			number := find_num(string(arr), i)
			for j := 0; j < number; j++ {
				new_str = append(new_str, arr[i-1])
			}
			isDigit = true
		} else if err == nil && i == 0 {
			return ""
		} else {
			isDigit = false
			new_str = append(new_str, value)
		}
	}
	return string(new_str)
}

func main() {
	str := "a4bc20d5e"

	new_str := unpacking_str(str)

	buf := "sa"

	buf = "mas" + buf
	fmt.Println(new_str, buf)
}
