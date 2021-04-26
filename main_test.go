package main

import "testing"

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2, 5)
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("David")
	}
}

func TestSayHello(t *testing.T) {
	t.Run("Test SayHello with parameter 'David' ", func(t *testing.T) {
		got := SayHello("David")
		want := "Hello, David"

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

func TestHelloHello(t *testing.T) {
	t.Run("testing Hello Hello", func(t *testing.T) {
		got := HelloHello()
		want := "Hello hello"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
