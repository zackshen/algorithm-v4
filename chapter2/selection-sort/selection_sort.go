package main

func swap(i, j int, data []int) {
	data[i], data[j] = data[j], data[i]
}

func SelectionSort(data []int) {
	N := len(data)
	for i := 0; i < N; i++ {
		minValue := data[i]
		minValueIndex := i
		// 从当前索引往后找，找到最小的一个数
		for j := i + 1; j < N; j++ {
			if data[j] <= minValue {
				minValue = data[j]
				minValueIndex = j
			}
		}
		// 将从索引右侧找到的最小数与当前数交换位置
		swap(i, minValueIndex, data)
	}
}

func main() {

}
