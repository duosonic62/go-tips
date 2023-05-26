package main

import (
	"fmt"
	"time"
)

func main() {
	t := "2021/12/23 22:05:10"

	parsed, err := parse(t)
	if err != nil {
		panic(err)
	}

	fmt.Printf(format(parsed))
}

func parse(s string) (*time.Time, error) {
	format := "2006/01/02 15:04:05"
	parsed, err := time.Parse(format, s)

	return &parsed, err
}

func format(time2 *time.Time) string {
	format := "2006-01-02 03:04:05"

	return time2.Format(format)
}