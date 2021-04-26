package main

import (
	"testing"
)

func TestHelloHello(t *testing.T) {
	got := HelloHello("Aziz")
	want := "Hello Aziz"

	t.Logf("Text: %s", want)

	if got != want {
		t.Errorf("Got %s, but want %s", got, want)
	}
}

func TestSayHello(t *testing.T) {
	t.Run("Test SayHello with parameter 'Aziz' ", func(t *testing.T) {
		got := SayHello("Aziz")
		want := "Hello, Aziz"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test SayHello if empty parameter", func(t *testing.T) {
		got := SayHello("")
		want := "Hello, friend"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test SayHello if empty parameter", func(t *testing.T) {
		got := SayHello("afista") // len 6
		want := "Hello, long name"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
