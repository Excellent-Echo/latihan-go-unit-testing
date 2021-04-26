package main

import (
	"testing"
)

// func TestLatihan2(t *testing.T) {
// 	fmt.Println(ChangeToIdr(100000))
// }

func TestLatihan2(t *testing.T) {
	t.Run("Test Latihan2 with parameter valid number ", func(t *testing.T) {
		got := ChangeToIdr(110000)
		want := "IDR 110,000,00"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
