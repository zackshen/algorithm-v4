package main

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

}
