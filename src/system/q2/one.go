package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("q2-1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s, %d, %f", "hoge", 1, 2.456)
}
