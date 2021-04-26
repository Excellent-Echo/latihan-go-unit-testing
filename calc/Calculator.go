package calc

import "errors"

func Calculator(num1 int, num2 int, method string) (int, error) {
	if num1 <= 0 || num2 <= 0 {
		return 0, errors.New("input must be valid number")
	} else {
		switch method {
		case "add":
			return num1 + num2, nil
		case "min":
			return num1 - num2, nil
		case "divide":
			return num1 / num2, nil
		case "multiple":
			return num1 * num2, nil
		default:
			return 0, errors.New("please choose method: add, min, divide, multiple")
		}
	}
}
