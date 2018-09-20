package main

func InsertSort(data []int) {
	N := len(data)
	for i := 1; i < N; i++ {
		value := data[i]
		index := i
		// 因为从左往右是从小到大的排列顺序
		// 从后往前找比value大的数
		// 如果比value大就往后挪一位
		// 直到比value小就退出循环
		for j := i; j > 0 && data[j-1] > value; j-- {
			data[j] = data[j-1]
			index = j - 1
		}
		// 将value放入最终停下的位置
		data[index] = value
	}
}

func main() {

}
