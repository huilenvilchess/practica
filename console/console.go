package main

import (
	"fmt"

	"github.com/mercadolibre/practica/calculadora"
)

func main() {
	opc := 0

	var numero1, numero2 int // Declarar variables

	for {
		fmt.Println("Elija una operación")
		fmt.Println("1-suma, 2-resta, 3-división, 4-multiplicación, Para salir presione A, Para continuar presione B")
		fmt.Scanf("%d", &opc)

		fmt.Println("1er número: ")
		fmt.Scanf("%d", &numero1)

		fmt.Println("2do número: ")
		fmt.Scanf("%d", &numero2)

		if opc == 1 {
			result, err := calculadora.Suma(numero1, numero2)
			if err != nil {
				fmt.Printf("Error %s", err.Error())
			} else {
				fmt.Printf("El resultado de la suma es %d", result)
			}
		} else if opc == 2 {
			result, err := calculadora.Resta(numero1, numero2)
			if err != nil {
				fmt.Printf("Error %s", err.Error())
			} else {
				fmt.Printf("El resultado de la resta es %d", result)
			}
		} else if opc == 3 {
			result, err := calculadora.Division(numero1, numero2)
			if err != nil {
				fmt.Printf("Error %s", err.Error())
			} else {
				fmt.Printf("El resultado de la division es %d", result)
			}
		} else if opc == 4 {
			result, err := calculadora.Multiplicacion(numero1, numero2)
			if err != nil {
				fmt.Printf("Error %s", err.Error())
			} else {
				fmt.Printf("El resultado de la multiplicacion es %d", result)
			}
		}

	}
}
