package main

import (
	"io"
	"os"
)

func main() {
	myIo := MyIO{}

	_, err := io.Copy(&myIo, os.Stdin)
	if err != nil {
		panic(err)
	}
}

type MyIO struct {
}

func (f *MyIO) Write(p []byte) (n int, err error) {

	return 0, nil
}