package main

func InsertSort(data []int) {
	N := len(data)
	for i := 1; i < N; i++ {
		value := data[i]
		index := i
		for j := i; j > 0 && data[j-1] > value; j-- {
			data[j] = data[j-1]
			index = j - 1
		}
		data[index] = value
	}
}

// 1 55 89

func main() {

}
