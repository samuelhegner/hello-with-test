package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("greet individual when name is supplied in english", func(t *testing.T) {
		got := Hello("Samuel", "english")
		want := "Hello Samuel!"

		assertCorrect(t, got, want)
	})

	t.Run("greet world on empty string in english", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World!"

		assertCorrect(t, got, want)
	})

	t.Run("greet individual when name is supplied in spanish", func(t *testing.T) {
		got := Hello("Samuel", "spanish")
		want := "Hola Samuel!"

		assertCorrect(t, got, want)
	})

	t.Run("greet world on empty string in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola Mundo!"

		assertCorrect(t, got, want)
	})

	t.Run("greet individual when name is supplied in french", func(t *testing.T) {
		got := Hello("Samuel", "french")
		want := "Bonjour Samuel!"

		assertCorrect(t, got, want)
	})

	t.Run("greet world on empty string in french", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour le monde!"

		assertCorrect(t, got, want)
	})

	t.Run("default to engish on unknown language", func(t *testing.T) {
		got := Hello("Sam", "FAKE_LANG")
		want := "Hello Sam!"

		assertCorrect(t, got, want)
	})
}

func assertCorrect(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
