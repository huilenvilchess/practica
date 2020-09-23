package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculadora_sumar_success(t *testing.T) {
	ass := assert.New(t)

	// When
	suma, err := Suma(10, 10)

	// Then
	ass.Equal(20, suma)
	ass.Nil(err)
}

func Test_calculadora_resta_success(t *testing.T) {
	ass := assert.New(t)

	// When
	resta, err := Resta(10, 10)

	// Then
	ass.Equal(0, resta)
	ass.Nil(err)
}
func Test_calculadora_division_success(t *testing.T) {
	ass := assert.New(t)

	// When
	division, err := Division(10, 10)

	// Then
	ass.Equal(1, division)
	ass.Nil(err)
}
func Test_calculadora_division_by_zero_should_return_error(t *testing.T) {
	ass := assert.New(t)

	// When
	_, err := Division(10, 0)

	// Then
	ass.NotNil(err)
}
func Test_calculadora_multiplicacion_success(t *testing.T) {
	ass := assert.New(t)

	// When
	multiplicacion, err := Multiplicacion(10, 10)

	// Then
	ass.Equal(100, multiplicacion)
	ass.Nil(err)
}
