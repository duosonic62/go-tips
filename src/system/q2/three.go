package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "json")
	sources := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Flush()
	writer := io.MultiWriter(gzipWriter, os.Stdout)

	encoder := json.NewEncoder(writer)
	encoder.Encode(sources)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
