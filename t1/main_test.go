package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Bob")
	want := "Hello, Bob"

	if got != want {
		t.Errorf("Hello(%q) = %q, want %q", "Bob", got, want)
	}
}
