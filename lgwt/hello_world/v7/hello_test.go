package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sunny", "")
		want := "Hello, Sunny"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// Testing for Spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sunny", "Spanish")
		want := "Hola, Sunny"
		assertCorrectMessage(t, got, want)
	})

	// Testing for French
	t.Run("in French", func(t *testing.T) {
		got := Hello("Sunny", "French")
		want := "Bonjour, Sunny"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
