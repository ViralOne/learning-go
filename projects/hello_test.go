package main

import (
	"testing"
)

func testH(t *testing.T) {
	got := Hello()
	want := "Hello, world."

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
