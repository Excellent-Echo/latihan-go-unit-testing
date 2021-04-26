package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Bantuan(data map[string]string) {
	jarakRumah, _ := strconv.Atoi(data["jarakRumah"])
	umur, _ := strconv.Atoi(data["umur"])

	if jarakRumah < 100 && data["berkeluarga"] == "ya" && umur > 20 {
		fmt.Println(data["name"] + " layak mendapat bantuan dari pemerintah")
	} else {
		fmt.Println(data["name"] + " tidak layak mendapat bantuan dari pemerintah")
	}
}

func TestLatihan4(t *testing.T) {
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

	Bantuan(data)
	Bantuan(data2)
	Bantuan(data3)
}
