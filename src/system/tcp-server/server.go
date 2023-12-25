package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func detectContentType(file os.FileInfo) (string, error) {
	switch {
	case strings.HasSuffix(file.Name(), ".html"):
		return "text/html", nil
	case strings.HasSuffix(file.Name(), ".json"):
		return "application/json", nil
	case strings.HasSuffix(file.Name(), ".txt"):
		return "text/plain", nil
	case strings.HasSuffix(file.Name(), ".jpeg") || strings.HasSuffix(file.Name(), ".jpg"):
		return "image/jpeg", nil
	}

	return "", fmt.Errorf("unknown file type")
}

func errorResponse(statusCode int, message string) *http.Response {
	return &http.Response{
		StatusCode:    statusCode,
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(strings.NewReader(message)),
		ContentLength: int64(len(message)),
	}
}

func handleRequest(request *http.Request) *http.Response {
	filePath := "./public" + request.URL.Path
	fmt.Println("Request " + filePath)

	// ファイルが存在しない場合
	if !exists(filePath) {
		log.Println("File not found: " + filePath)
		return errorResponse(404, "File not found: "+filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return errorResponse(500, "Internal Server Error")
	}

	fileInfo, err := file.Stat()
	contentType, err := detectContentType(fileInfo)
	if err != nil {
		log.Println(err)
		return errorResponse(500, "Internal Server Error")
	}

	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        make(http.Header),
		Body:          file,
		ContentLength: fileInfo.Size(),
	}
	response.Header.Set("Content-Type", contentType)
	return response
}

func handleSession(conn net.Conn) {
	fmt.Println("Client connected")
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		response := handleRequest(request)

		response.Write(conn)
	}

}

func main() {
	ln, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			handleSession(conn)
		}()
	}
}
