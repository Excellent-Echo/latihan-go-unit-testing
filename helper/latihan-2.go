package helper

import (
	"errors"
	"fmt"
)

func validateChangeToIdr(price int) (bool, error) {
	if price == 0 {
		return false, errors.New("tidak boleh tanpa value")
	}
	return true, nil
}

func catchChangeToIdr() {
	if r := recover(); r != nil {
		fmt.Println("Keterangan error :", r)
	} else {
		fmt.Println("Aplikasi berjalan sempurna")
	}
}

func ChangeToIdr(price int) string {
	defer catchChangeToIdr()

	var hasil string
	if valid, err := validateChangeToIdr(price); valid {
		if price < 0 {
			return "-" + ChangeToIdr(-price)
		}
		if price < 1000 {
			return fmt.Sprintf("%d", price)
		}
		result := ChangeToIdr(price/1000) + "," + fmt.Sprintf("%03d", price%1000)
		//final := "IDR " + result + "," + "00"
		//fmt.Println("Hasil", final)
		hasil = result
	} else {
		panic(err.Error())
		fmt.Println("err")
	}
	return hasil

	//s := strconv.Itoa(number)
	//firstDigit := string(s[0]) + string((s[1])) + string(s[2])
	//secondDigit := string(s[3]) + string(s[4]) + string(s[5])
	//
	//values := []string{}
	//values = append(values, firstDigit)
	//values = append(values, secondDigit)
	//fmt.Println(values)
	//
	//
	//result := strings.Join(values, ",")
	//fmt.Println("IDR " + result + "," + "00")

}
