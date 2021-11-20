package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Say hello world when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("John", "spanish")
		want := "Hola, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("Adam", "french")
		want := "Bonjour, Adam"
		assertCorrectMessage(t, got, want)
	})
}
