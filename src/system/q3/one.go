package main

import (
	"io"
	"os"
)

func main() {
	old, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer old.Close()

	file, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, old)
}
