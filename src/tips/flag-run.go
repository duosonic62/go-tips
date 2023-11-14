package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	max := commandLine.Int("max", 255, "max value")
	name := commandLine.String("name", "", "my name")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
	}

	if *max > 999 {
		fmt.Fprintf(os.Stderr, "invalid max value: %v\n", *max)
		return 1
	}

	if *name == "" {
		fmt.Fprintf(os.Stderr, "name must be provided")
		return 1
	}

	return 0
}
