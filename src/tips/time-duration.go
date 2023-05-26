package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("3s")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	d = time.Minute * 5
	fmt.Println(d)

	n := time.Now()
	fmt.Println(n)

	fmt.Println(n.Add(d))
}
