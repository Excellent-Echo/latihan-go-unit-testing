package main

import (
	"fmt"
	"goAdvance/helper"
)

func main() {
	// Latihan 1
	fmt.Println(helper.Total(110000))
	fmt.Println(helper.Total(220000))
	fmt.Println(helper.Total(500000))
	fmt.Println(helper.Total(1000000))
	fmt.Println(helper.Total(0))
	fmt.Println("-----------------------------------")

	//	latihan 2
	fmt.Println(helper.ChangeToIdr(100000))
	fmt.Println(helper.ChangeToIdr(1200000))
	fmt.Println(helper.ChangeToIdr(140350000))
	fmt.Println(helper.ChangeToIdr(0))

	fmt.Println("-----------------------------------")

	// latihan 3
	fmt.Println(helper.GenapGanjil(1, 2, 3, 4, 5))
	fmt.Println(helper.GenapGanjil(4, 2))
	fmt.Println(helper.GenapGanjil(10, 20, 30, 13))
	fmt.Println(helper.GenapGanjil(30, 13, 13, 77, 33, 55, 17, 13))
	fmt.Println(helper.GenapGanjil())

	fmt.Println("-----------------------------------")

	// latihan 4
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
	var data3 = map[string]string{
		"name":        "budi",
		"umur":        "45",
		"jarakRumah":  "120",
		"berkeluarga": "ya",
	}
	fmt.Println(helper.GovermentHelper(data))
	fmt.Println(helper.GovermentHelper(data2))
	fmt.Println(helper.GovermentHelper(data3))
	fmt.Println("-----------------------------------")

	// latihan 5
	helper.ChangeNumtoStr(100000)
	helper.ChangeNumtoStr(111000)
	helper.ChangeNumtoStr(5124000)
	helper.ChangeNumtoStr(1543)
	helper.ChangeNumtoStr(1234678)
	helper.ChangeNumtoStr(0)

	fmt.Println("-----------------------------------")

	// latihan 6
	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}

	fmt.Println(helper.TotalPajak(belanja1))

}
