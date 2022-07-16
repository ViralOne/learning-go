package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("testHello", func(t *testing.T) {
		got := Hello("Milu")
		want := "Hello, Milu"
		assertCorrectMessage(t, got, want)
	})

	// failing test
	t.Run("Test call empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, there"
		assertCorrectMessage(t, got, want)
	})
}
