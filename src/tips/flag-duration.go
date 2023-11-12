package main

import (
	"flag"
	"time"
)

// ./main -sleep 3s
func main() {
	var sleep time.Duration
	flag.DurationVar(&sleep, "sleep", time.Second, "sleep time")
	flag.Parse()

	time.Sleep(sleep)
}
