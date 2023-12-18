package main

import (
	"archive/zip"
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("q3-3.zip")
	if err != nil {
		panic(err)
	}

	zipWrite := zip.NewWriter(file)
	defer zipWrite.Close()

	writer, err := zipWrite.Create("q3-3.txt")
	if err != nil {
		panic(err)
	}

	io.CopyN(writer, rand.Reader, 1024)
}
