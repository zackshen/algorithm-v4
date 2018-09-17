package main

import (
	"fmt"
)

func swap(i, j int, data []int) {
	data[i], data[j] = data[j], data[i]
}

func SelectionSort(data []int) {
	N := len(data)
	for i := 0; i < N; i++ {
		minValue := data[i]
		minValueIndex := i
		for j := i + 1; j < N; j++ {
			if data[j] <= minValue {
				minValue = data[j]
				minValueIndex = j
			}
		}
		swap(i, minValueIndex, data)
	}
}

func main() {
	data := []int{55, 1, 89, 34, 46, 27, 86, 94, 5, 6}
	SelectionSort(data)
	fmt.Printf("%v\n", data)

	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	SelectionSort(data2)
	fmt.Printf("%v\n", data2)

}
