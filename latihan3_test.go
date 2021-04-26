package main

import (
	"fmt"
	"testing"
)

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

func TestLatihan3(t *testing.T) {
	fmt.Println(GenapGanjil(1, 2, 3, 4, 5))
}
