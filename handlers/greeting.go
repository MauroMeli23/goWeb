package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name     string
	LastName string
}

func CreateGreetings(c *gin.Context) {
	// Parsear el cuerpo JSON de la solicitud en CreateGreetings
	var req Person
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	greeting := "Hola " + req.Name + " " + req.LastName
	c.JSON(http.StatusCreated, greeting)
}
