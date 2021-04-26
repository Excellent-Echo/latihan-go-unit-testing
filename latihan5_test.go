package main

import (
	"fmt"
	"strings"
	"testing"
)

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
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", puluhan[tens], satuan[units])
				hasil = append(hasil, word)
			} else {
				hasil = append(hasil, puluhan[tens])
			}
			break
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

func TestLatihan5(t *testing.T) {
	fmt.Println(ChangeNumtoStr(100000))
	fmt.Println(ChangeNumtoStr(111000))
	fmt.Println(ChangeNumtoStr(5124000))
	fmt.Println(ChangeNumtoStr(1543))
	fmt.Println(ChangeNumtoStr(1234678))
	fmt.Println(ChangeNumtoStr(0))
}
