package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	stream, pipeWriter := io.Pipe()
	go func() {
		defer pipeWriter.Close()

		// a
		programming.Seek(5, 0)
		io.CopyN(pipeWriter, programming, 1)

		// s
		io.CopyN(pipeWriter, system, 1)

		// c
		io.CopyN(pipeWriter, computer, 1)

		// i
		programming.Seek(8, 0)
		io.CopyN(pipeWriter, programming, 1)

		// i
		programming.Seek(8, 0)
		io.CopyN(pipeWriter, programming, 1)
	}()

	// ASCII
	io.Copy(os.Stdout, stream)
}
