package main

import (
	"github.com/MauroMeli23/goWeb/handlers"
	"github.com/MauroMeli23/goWeb/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	// Crea una instancia del motor Gin
	r := gin.Default()

	var products, err = utils.LoadProducts()
	if err != nil {
		log.Fatal("Error al cargar los productos:", err)
	}

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

	//GetAllProducts
	r.GET("/products", func(c *gin.Context) {
		handlers.GetAllProducts(c, products)
	})

	//GetProductByID
	r.GET("/products/:id", func(c *gin.Context) {
		handlers.GetProductByID(c, products)
	})

	//SearchProduct
	r.GET("/products/search", func(c *gin.Context) {
		handlers.GetProductByName(c, products)
	})

	// Ejecuta el servidor en el puerto 8080
	r.Run(":8080")
}
