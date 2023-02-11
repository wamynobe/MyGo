package main

func main() {
	print(HardMinSwaps([]int{4, 3, 1, 2}))
}

func HardMinSwaps(arr []int) int {
	count := 0
	if len(arr) <= 1 {
		return len(arr)
	}
	for i := 0; i < len(arr); i++ {
		for arr[i] != i+1 {
			tmp := arr[i]
			arr[i] = arr[tmp-1]
			arr[tmp-1] = tmp
			count++
		}
	}
	return count
}
