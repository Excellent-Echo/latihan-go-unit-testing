package main

import (
	"fmt"
	"testing"
)

func Belanja(total int) (ppn int, beforePpn int) {
	beforePpn = total * 10 / 11
	ppn = total / 11
	return
}
func TestLatihan1(t *testing.T) {
	pajak, harga := Belanja(110000)
	fmt.Printf("Pajak: %v, Harga: %v\n", pajak, harga)
}
