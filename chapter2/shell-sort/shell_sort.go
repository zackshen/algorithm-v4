package main

func swap(i, j int, data []int) {
	data[i], data[j] = data[j], data[i]
}

func ShellSort(data []int) {
	N := len(data)
	h := 1
	for {
		if h >= N/3 {
			break
		}
		h = 3*h + 1
	}

	for {
		if h < 1 {
			break
		}

		for i := h; i < N; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				swap(j, j-h, data)
			}
		}

		h = h / 3
	}
}

func main() {

}
