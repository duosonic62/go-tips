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

func TestCreateProfile2(t *testing.T) {
	t.Setenv("DATABASE_URL", "127.0.0.1")
	_, err := CreateProfile("test")
	if err != nil {
		t.Fatal(err)
	}
}
