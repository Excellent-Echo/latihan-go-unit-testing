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

func ChangeToIdr(nominal int) string {
	idr := strconv.Itoa(nominal)
	thousand := 3
	if nominal < 0 {
		thousand++
	}
	for i := len(idr); i > thousand; {
		i -= 3
		idr = idr[:i] + "," + idr[i:]
	}
	return "IDR " + idr + ",00"
}
