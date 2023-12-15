package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("q2-2.txt")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"hoge", "1", "2.456"})
	writer.Write([]string{"huga", "2", "3.456"})

}
