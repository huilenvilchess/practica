package calculadora

import (
"errors"
)

func Suma(a, b int) (int, error) {
	return a + b, nil
}

func Resta(a, b int) (int, error) {
	return a - b, nil
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func Multiplicacion(a, b int) (int, error) {
	return a * b, nil
}
