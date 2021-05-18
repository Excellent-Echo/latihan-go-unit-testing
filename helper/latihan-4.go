package helper

import (
	"strconv"
)

func GovermentHelper(data map[string]string) string {
	var result string

	jarakRumah, _ := strconv.Atoi(data["jarakRumah"])
	umur, _ := strconv.Atoi(data["umur"])

	if jarakRumah < 100 && data["berkeluarga"] == "ya" && umur > 20 {
		result = data["name"] + " layak mendapat bantuan dari pemerintah"
	} else {
		result = data["name"] + " tidak layak mendapat bantuan dari pemerintah"
	}
	return result
}
