package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func SayHello(name string) string {
	if name == "" {
		return "Hello, friend"
	} else if len(name) > 5 {
		return "Hello, long name"
	} else {
		return "Hello, " + name
	}
}

func HelloHello() string {
	return "Hello hello"
}

func Calculate(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	// 3rd party

	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
	}

	fmt.Println(ac.FormatMoney(100000))

	// result, err := calc.Calculator(1, 2, "add")

	// greeting := hello.Hello()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(greeting)
	// fmt.Println(result)
}
