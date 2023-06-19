package main

import "fmt"

func main() {
	var n I = 0

	n = n.Add(1).Add(2)
	fmt.Println(n)

	add := n.Add
	fmt.Println(add(3))

	fmt.Printf("%T\n", n.Add) // func(int) main.I
	fmt.Printf("%T\n", I.Add) // func(main.I, int) main.I

	n = 1
	fmt.Println(n.Add(2))
	fmt.Println(I.Add(n, 2))
}

type I int

func (i I) Add(n int) I {
	return i + I(n)
}
