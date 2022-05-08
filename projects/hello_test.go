package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("testHello", func(t *testing.T) {
		got := Hello("Milu")
		want := "Hello, Milu"

		if got != want {
			t.Errorf("Hello() = %q, want %q", got, want)
		}
	})

	// failing test
	t.Run("Test call empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, there"

		if got != want {
			t.Errorf("Hello() = %q, want %q", got, want)
		}
	})
}
