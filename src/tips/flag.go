package main

import "flag"

func main() {
	// return pointer
	max := flag.Int("max", 255, "max value")
	name := flag.String("name", "something", "my name")

	var min int
	var id string
	flag.IntVar(&min, "min", 0, "min value")
	flag.StringVar(&id, "id", "id", "your account id")

	flag.Parse()

	println(*name, *max, min, id)
}
