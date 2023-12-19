package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "application/zip")
	w.Header().Set("Content-Type", "attachment; filename=q3-4.zip")
	sources := "Hello, World"

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("q3-4.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(writer, strings.NewReader(sources))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
