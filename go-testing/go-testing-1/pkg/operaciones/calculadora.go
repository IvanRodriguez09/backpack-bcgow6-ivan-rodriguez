package operations

import "errors"

func substract(num1, num2 int) (int, error) {
	return num1 - num2, nil
}

func divide(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("el denominador no puede ser 0")
	}
	return float64(num1 / num2), nil
}
