package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

func main() {

}

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func decodeJson(f *os.File, conn io.Reader) {
	var data Data
	dec := json.NewDecoder(f)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		scanner.Text()
		err := dec.Decode(&data)
		if err != nil {
			break
		}
	}
}
