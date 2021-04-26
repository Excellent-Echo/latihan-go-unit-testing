package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Ppn(data map[string]int) string {
	sum := 0

	for _, val := range data {
		ppn := val / 11
		sum += ppn
	}
	return strconv.Itoa(sum)
}

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

	totalPpn := Ppn(belanja1)
	fmt.Println("total PPN yang diterima sebesar " + totalPpn)
}
