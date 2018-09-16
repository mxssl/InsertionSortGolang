package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 6, 7, 8, 5, 4, 3, 2, 1}
	fmt.Println(insertionSort(arr))
}

func insertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			swap(j, j-1, array)
		}
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}
