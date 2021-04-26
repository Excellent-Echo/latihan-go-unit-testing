package main

import (
	"testing"
)

func TestLatihan3(t *testing.T) {
	t.Run("Test Latihan3 case ganjil ", func(t *testing.T) {
		got := GenapGanjil(1, 2, 3, 4, 5)
		want := "angka terbanyak adalah ganjil"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test Latihan2 case genap ", func(t *testing.T) {
		got := GenapGanjil(4, 2)
		want := "angka terbanyak adalah genap"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test Latihan2 case no number ", func(t *testing.T) {
		got := GenapGanjil()
		want := "tidak ada angka"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

}
