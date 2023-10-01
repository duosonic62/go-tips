package main

import "testing"

// run command
// go test -fuzz FuzzDoSomething -fuzztime 10s
func FuzzDoSomething(f *testing.F) {
	f.Add("test&&&")
	f.Fuzz(func(f *testing.T, s string) {
		DoSomething(s)
	})
}
