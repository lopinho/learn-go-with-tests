package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Saying 'hello world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Andre", "")
		want := "Hello, Andre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Andre", SPANISH)
		want := "Hola, Andre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in french", func(t *testing.T) {
		got := Hello("Andre", FRENCH)
		want := "Bonjour, Andre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in portuguese", func(t *testing.T) {
		got := Hello("Andre", PORTUGUESE)
		want := "Olá, Andre"
		assertCorrectMessage(t, got, want)
	})
}
