package main

import "testing"

func TestHello(t *testing.T) {
	// t testing.TB:
	// This is an interface that *testing.T and *testing.B both satisfy.
	// This means that you can call this helper function from both a test, or a benchmark.
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper() tells the test suite that this method is a helper.
		// By doing this, when it fails the line number reported will be the function call rather than the helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Frenchie", "French")
		want := "Bonjour, Frenchie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("公子", "Japanese")
		want := "ご機嫌よう、公子"
		assertCorrectMessage(t, got, want)
	})
}
