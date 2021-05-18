package helper

import (
	"strconv"
)

func TotalPajak(belanja1 map[string]int) string {
	var result string

	total := 0
	for _, value := range belanja1 {
		//fmt.Println(value)
		pajak := value / 11
		//fmt.Println(pajak)
		total += pajak
	}

	toString := strconv.Itoa(total)
	result = "total PPN yang diterima sebesar " + toString
	return result

}
