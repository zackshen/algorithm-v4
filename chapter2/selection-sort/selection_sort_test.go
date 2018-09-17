package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := []int{55, 1, 89, 34, 46, 27, 86, 94, 5, 6}
	SelectionSort(data)
	assert.Equal(t, data, []int{1, 5, 6, 27, 34, 46, 55, 86, 89, 94}, "they should be equal")

	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	SelectionSort(data2)
	assert.Equal(t, data2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "they should be equal")
}
