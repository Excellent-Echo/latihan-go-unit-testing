package main

import (
	"testing"
)

// func TestLatihan5(t *testing.T) {
// 	fmt.Println(ChangeNumtoStr(100000))
// 	fmt.Println(ChangeNumtoStr(111000))
// 	fmt.Println(ChangeNumtoStr(5124000))
// 	fmt.Println(ChangeNumtoStr(1543))
// 	fmt.Println(ChangeNumtoStr(1234678))
// 	fmt.Println(ChangeNumtoStr(0))
// }

func TestLatihan5(t *testing.T) {
	t.Run("Test Latihan5 with parameter valid number ", func(t *testing.T) {
		got := ChangeNumtoStr(111000)
		want := "seratus sebelas ribu"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}
