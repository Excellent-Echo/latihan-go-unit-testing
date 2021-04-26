package main

import (
	"testing"
)

func TestLatihan4(t *testing.T) {
	var data = map[string]string{
		"name":        "andi",
		"umur":        "30",
		"jarakRumah":  "50",
		"berkeluarga": "ya",
	}

	var data2 = map[string]string{
		"name":        "santi",
		"umur":        "19",
		"jarakRumah":  "80",
		"berkeluarga": "ya",
	}

	t.Run("Test Latihan4 case layak mendapat bantuan ", func(t *testing.T) {
		got := Bantuan(data)
		want := "andi layak mendapat bantuan dari pemerintah"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test Latihan4 case tidak layak mendapat bantuan ", func(t *testing.T) {
		got := Bantuan(data2)
		want := "santi tidak layak mendapat bantuan dari pemerintah"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
