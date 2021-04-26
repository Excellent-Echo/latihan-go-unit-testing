package main

import (
	"fmt"
	"strconv"
	"strings"
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

func groupOfThreeArray(number int) []int {
	array := []int{}
	for number > 0 {
		array = append(array, number%1000)
		number = number / 1000
		// fmt.Printf("%v %v\n", array, number)
	}
	return array
}

func ChangeNumtoStr(input int) string {
	satuan := [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	belasan := [...]string{"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}
	puluhan := [...]string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}
	ribuan := [...]string{"", "ribu", "juta", "milyar", "triliun"}
	hasil := []string{}
	if input < 0 {
		hasil = append(hasil, "minus")
		input *= -1
	}
	// ubah inputan ke array triplets
	triplets := groupOfThreeArray(input)
	if len(triplets) == 0 {
		return "nol"
	}
	for i := len(triplets) - 1; i >= 0; i-- {
		triplet := triplets[i]
		// fmt.Printf("Triplet: %d (i=%d)\n", triplet, i)
		if triplet == 0 {
			continue
		}
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		// fmt.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds == 1 {
			hasil = append(hasil, "seratus")
		} else if hundreds > 0 {
			hasil = append(hasil, satuan[hundreds], "ratus")
		}
		switch tens {
		case 0:
			hasil = append(hasil, satuan[units])
		case 1:
			hasil = append(hasil, belasan[units])
			// break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", puluhan[tens], satuan[units])
				hasil = append(hasil, word)
			} else {
				hasil = append(hasil, puluhan[tens])
			}
			// break
		}
		mega := ribuan[i]
		if mega != "" {
			if i == 1 && triplet == 1 {
				hasil = append(hasil[0:len(hasil)-1], "seribu")
			} else {
				hasil = append(hasil, mega)
			}
		}
	}
	return strings.Join(hasil, " ")
}
