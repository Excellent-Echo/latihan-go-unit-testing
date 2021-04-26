package helper

import (
	"errors"
	"fmt"
)

func validateGanjilGenap(genap int, ganjil int) (bool, error) {
	if genap == 0 && ganjil == 0 {
		return false, errors.New("tidak boleh kosong")
	}
	return true, nil
}

func catchGanjilGenap() {
	if r := recover(); r != nil {
		fmt.Println("keterangan error :", r)
	} else {
		fmt.Println("Aplikasi berjalan sempurna")
	}
}

func GenapGanjil(number ...int) string {
	defer catchGanjilGenap()

	var (
		genap, ganjil int
		hasil         string
	)
	for _, value := range number {
		if value%2 == 0 {
			genap++
		}
		if value%2 == 1 {
			ganjil++
		}
	}
	if valid, err := validateGanjilGenap(genap, ganjil); valid {
		if genap > ganjil {
			hasil = "Angka terbanyak adalah genap"
		}
		if ganjil > genap {
			hasil = "Angka terbanyak adalah ganjil"
		}
	} else {
		panic(err.Error())
		fmt.Println("err")
	}

	return hasil

}
