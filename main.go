package main

import "strconv"

func Belanja(total int) string {
	beforePpn := total * 10 / 11
	ppn := total / 11

	if total <= 0 {
		return "Input must be a valid number"
	} else {
		return "Pajak: " + strconv.Itoa(ppn) + ", Harga: " + strconv.Itoa(beforePpn)
	}
}
