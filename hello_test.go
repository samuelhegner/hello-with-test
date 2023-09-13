package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Samuel")
	want := "Hello Samuel!"

	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
