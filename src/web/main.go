package main

import (
	"fmt"
	"net/http"
)

func myHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		fmt.Fprintf(writer, "Get:  Hello, World!")
	case http.MethodPost:
		fmt.Fprintf(writer, "Post: Hello, World!")
	}
}

func main() {
	http.HandleFunc("/", myHandler)

	http.ListenAndServe(":8080", nil)
}
