package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	sortMapByKey(m, func(v string) {
		fmt.Printf("v: %s\n", v)
	})
}

func sortMapByKey(m map[string]int, do func(v string)) {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}

	sort.Strings(keys)
	for _, key := range keys {
		do(key)
	}
}
