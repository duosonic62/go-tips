package main

import "fmt"

func main() {
	m := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	find(m, "John", func(v int) {
		fmt.Printf("v: %d\n", v)
	})

	// do not print
	find(m, "Johnny", func(v int) {
		fmt.Printf("v: %d\n", v)
	})
}

func find(m map[string]int, key string, do func(v int)) {
	if v, ok := m[key]; ok {
		do(v)
	}
}