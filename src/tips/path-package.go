package main

import (
	"fmt"
	"path/filepath"
)

func main() {

}

func filePath() {
	current := filepath.Dir("./")
	abs, err := filepath.Abs(current)
	if err != nil {
		panic(err)
	}

	fmt.Println(current)
	fmt.Println(abs)
}
