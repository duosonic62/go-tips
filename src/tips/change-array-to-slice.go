package main

func main() {
	changeArrayToSlice([3]int{1, 2, 3})
}

func changeArrayToSlice(arr [3]int) []int {
	return arr[:]
}
