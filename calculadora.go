package main

type calculadora struct {
}

func (c calculadora) Sumar(a, b int) int {
	return a + b
}

func (c calculadora) Resta(a, b int) int {
	return a - b
}

func (c calculadora) Division(a, b int) int {
	return a / b
}

func (c calculadora) Multiplicacion(a, b int) int {
	return a * b
}
