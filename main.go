package main

import (
	"github.com/MauroMeli23/goWeb/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Crea una instancia del motor Gin
	r := gin.Default()

	//Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	//Ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//Saludo
	r.POST("/saludo", handlers.CreateGreetings)

	// Ejecuta el servidor en el puerto 8080
	r.Run(":8080")
}
