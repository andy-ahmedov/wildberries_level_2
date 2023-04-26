package main

import "fmt"

func reverse(slice []int, newArr []int) []int {
	newArr = append(newArr, slice[len(slice)-1])
	if len(slice) == 1 {
		return newArr
	}
	return (reverse(slice[:len(slice)-1], newArr))
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	newArr := []int{}

	newArr = reverse(arr, newArr)
	fmt.Println(newArr)
}
