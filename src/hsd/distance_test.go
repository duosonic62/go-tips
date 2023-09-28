package main

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestStringDistance_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{name: "lhs is longer than rhs", lhs: "foo", rhs: "fo", want: -1},
		{name: "lhs is shorter than rhs", lhs: "fo", rhs: "foo", want: -1},
		{name: "no diff", lhs: "foo", rhs: "foo", want: 0},
		{name: "3 diffs", lhs: "foo", rhs: "hhh", want: 3},
		{name: "multibyte", lhs: "あいう", rhs: "あいえ", want: 1},
	}

	for _, tc := range tests {
		got := StringDistance(tc.lhs, tc.rhs)
		if got != tc.want {
			t.Fatalf("%s expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestStringDistance_Skip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skipf("this test is skipped")
	}

	// this test is skipped if it runs with short option like `go test -short`
	if testing.Short() {
		t.SkipNow()
	}
}

func TestMain(m *testing.M) {
	// setup
	log.Println("Before")
	ret := m.Run()

	// teardown
	log.Println("After")
	os.Exit(ret)
}
