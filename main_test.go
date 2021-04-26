// package main

// import "testing"

// func TestSayHello(t *testing.T) {
// 	t.Run("Test SayHello with parameter 'David' ", func(t *testing.T) {
// 		got := SayHello("David")
// 		want := "Hello, David"

// 		t.Logf("Test : %s", want)

// 		if got != want {
// 			t.Errorf("got %s, but want %s", got, want)
// 		}
// 	})

// 	t.Run("Test SayHello if empty parameter", func(t *testing.T) {
// 		got := SayHello("")
// 		want := "Hello, friend"

// 		t.Logf("Test : %s", want)

// 		if got != want {
// 			t.Errorf("got %s, but want %s", got, want)
// 		}
// 	})

// 	t.Run("Test SayHello if empty parameter", func(t *testing.T) {
// 		got := SayHello("afista") // len 6
// 		want := "Hello, long name"

// 		t.Logf("Test : %s", want)

// 		if got != want {
// 			t.Errorf("got %s, but want %s", got, want)
// 		}
// 	})
// }
