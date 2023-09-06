package handlers

import (
	"github.com/MauroMeli23/goWeb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllProducts(c *gin.Context, products []utils.Product) {
	// Usar la variable global 'products' para responder con los productos
	c.JSON(http.StatusCreated, products)
}

func GetProductByID(c *gin.Context, products []utils.Product) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto no válido"})
		return
	}
	if productID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	var product utils.Product
	for _, p := range products {
		if p.ID == productID {
			product = p
			break
		}
	}
	c.JSON(http.StatusCreated, product)
}

func GetProductByName(c *gin.Context, products []utils.Product) {
	paramValue := c.Query("name")

	if paramValue == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Debe ingresar un nombre"})
		return
	}
	var product utils.Product
	for _, p := range products {
		if p.Name == paramValue {
			product = p
			break
		}

	}
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}
