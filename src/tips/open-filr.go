package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	_, err = fmt.Fprintf(f, "write to file")
	if err != nil {
		panic(err)
	}
}
