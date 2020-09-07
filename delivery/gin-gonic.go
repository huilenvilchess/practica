package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func suma() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/suma", func(c *gin.Context) {
		var payload CalculadoraPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		suma := payload.Numero1 + payload.Numero2
		result := Result{
			Resultado: suma,
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run()

}

func resta() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/resta", func(c *gin.Context) {
		var payload CalculadoraPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resta := payload.Numero1 - payload.Numero2
		result := Result{
			Resultado: resta,
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run()

}

func division() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/division", func(c *gin.Context) {
		var payload CalculadoraPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resta := payload.Numero1 / payload.Numero2
		result := Result{
			Resultado: division,
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run()

}
func multiplicacion() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/multiplicacion", func(c *gin.Context) {
		var payload CalculadoraPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resta := payload.Numero1 * payload.Numero2
		result := Result{
			Resultado: multiplicacion,
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run()

}

// Binding from JSON
type CalculadoraPayload struct {
	Numero1 int `json:"numero1"`
	Numero2 int `json:"numero2"`
}
type Result struct {
	Resultado int `json:"resultado"`
}
