package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculadora_sumar(t *testing.T) {
	ass := assert.New(t)

	// Given
	calc := &calculadora{}

	// When
	suma := calc.Sumar(10, 10)

	// Then
	ass.Equal(20, suma)
}

func Test_calculadora_resta(t *testing.T) {
	ass := assert.New(t)

	// Given
	calc := &calculadora{}

	// When
	resta := calc.Resta(10, 10)

	// Then
	ass.Equal(0, resta)

}

func Test_calculadora_division(t *testing.T) {
	ass := assert.New(t)

	// Given
	calc := &calculadora{}

	// When
	division := calc.Division(10, 10)

	// Then
	ass.Equal(1, division)

}

func Test_calculadora_multiplicacion(t *testing.T) {
	ass := assert.New(t)

	// Given
	calc := &calculadora{}

	// When
	multiplicacion := calc.Multiplicacion(10, 10)

	// Then
	ass.Equal(100, multiplicacion)

}
