package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("q3-1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	randomReader := rand.Reader
	io.CopyN(file, randomReader, 1024)
}
