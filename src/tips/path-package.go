package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
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

func file() {
	files := []string{}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	fmt.Println(files)
}