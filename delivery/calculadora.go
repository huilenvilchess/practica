package main

import (
	"errors"
)

type calculadora struct {
}

func (c calculadora) Suma(a, b int) (int, error) {
	return a + b, nil
}

func (c calculadora) Resta(a, b int) (int, error) {
	return a - b, nil
}

func (c calculadora) Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func (c calculadora) Multiplicacion(a, b int) (int, error) {
	return a * b, nil
}
