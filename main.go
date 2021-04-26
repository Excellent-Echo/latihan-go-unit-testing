package main

import (
	"strconv"
)

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

func GenapGanjil(numbers ...int) string {
	countGanjil := 0
	countGenap := 0
	if len(numbers) <= 0 {
		return "tidak ada angka"
	}
	for _, num := range numbers {
		if num%2 == 0 {
			countGenap++
		} else {
			countGanjil++
		}
	}
	if countGanjil > countGenap {
		return "angka terbanyak adalah ganjil"
	} else if countGenap > countGanjil {
		return "angka terbanyak adalah genap"
	}

	return ""
}

func Bantuan(data map[string]string) string {
	jarakRumah, _ := strconv.Atoi(data["jarakRumah"])
	umur, _ := strconv.Atoi(data["umur"])

	if jarakRumah < 100 && data["berkeluarga"] == "ya" && umur > 20 {
		return data["name"] + " layak mendapat bantuan dari pemerintah"
	} else {
		return data["name"] + " tidak layak mendapat bantuan dari pemerintah"
	}
}
