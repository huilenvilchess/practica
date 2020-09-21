package main

import (
	"github.com/mercadolibre/practica/calculadora"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/suma", func(c *gin.Context) {
		ejecutarOperacion(c, calculadora.Suma)
	})

	r.POST("/resta", func(c *gin.Context) {
		ejecutarOperacion(c, calculadora.Resta)
	})

	r.POST("/division", func(c *gin.Context) {
		ejecutarOperacion(c, calculadora.Division)
	})

	r.POST("/multiplicacion", func(c *gin.Context) {
		ejecutarOperacion(c, calculadora.Multiplicacion)
	})

	r.Run()
}

func ejecutarOperacion(c *gin.Context, operacion func(a, b int) (int, error)) {
	var payload CalculadoraPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado, err := operacion(payload.Numero1, payload.Numero2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result := Result{
		Resultado: resultado,
	}

	c.JSON(http.StatusOK, result)
}

// Binding from JSON
type CalculadoraPayload struct {
	Numero1 int `json:"numero1"`
	Numero2 int `json:"numero2"`
}
type Result struct {
	Resultado int `json:"resultado"`
}
