package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	source := strings.NewReader("Too long text.Too long text.Too long text.Too long text.Too long text.Too long text.Too long text.")

	copyN(os.Stdout, source, 16)
	println("---")
	io.CopyN(os.Stdout, source, 16)
}

func copyN(dest io.Writer, src io.Reader, length int) {
	lReade := io.LimitReader(src, int64(length))
	io.Copy(dest, lReade)
}
