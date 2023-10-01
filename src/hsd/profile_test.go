package main

import (
	"path/filepath"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	dir := t.TempDir() // create temp dir
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)
	if err != nil {
		t.Fatal(err)
	}

	want := true
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
