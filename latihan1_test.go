package main

import (
	"testing"
)

func TestLatihan1(t *testing.T) {
	t.Run("Test Latihan1 with parameter valid number ", func(t *testing.T) {
		got := Belanja(110000)
		want := "Pajak: 10000, Harga: 100000"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test Latihan1 with parameter unvalid number ", func(t *testing.T) {
		got := Belanja(0)
		want := "Input must be a valid number"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

}
