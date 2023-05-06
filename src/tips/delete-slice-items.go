package main

import "fmt"

func main() {
	useAppend()
	useCopy()
}

func useAppend() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a = append(a[:3], a[3+1:]...)
	fmt.Printf("%d\n", a)
}

func useCopy() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a = a[:3+copy(a[3:], a[3+1:])]
	fmt.Printf("%d\n", a)
}
