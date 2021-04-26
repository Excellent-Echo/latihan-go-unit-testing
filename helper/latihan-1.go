package helper

import (
	"errors"
	"fmt"
)

func validate(price int) (bool, error) {
	if price == 0 {
		return false, errors.New("tidak boleh tanpa value")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Keterangan error :", r)
	} else {
		fmt.Println("Aplikasi berjalan sempurna")
	}
}

func Total(price int) int {
	defer catch()

	var (
		pajak, normalPrice int
	)

	if valid, err := validate(price); valid {
		pajak = price / 11
		normalPrice = price - pajak
	} else {
		panic(err.Error())
		fmt.Println("end")
	}

	fmt.Println("harga normal :", normalPrice)
	fmt.Println("Pajak PPN :", pajak)
	//fmt.Println("Total PPN :", price)

	return pajak
}
