package main

import (
	"testing"
)

func TestLatihan6(t *testing.T) {
	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}
	t.Run("Test Latihan5 with parameter valid number ", func(t *testing.T) {
		got := Ppn(belanja1)
		want := "total PPN yang diterima sebesar 472000"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
