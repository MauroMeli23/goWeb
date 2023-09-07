package handlers

import (
	"github.com/MauroMeli23/goWeb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

func AddNewProduct(c *gin.Context, products *[]utils.Product) {
	var newProduct utils.Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newProduct.Name == "" || newProduct.CodeValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name y CodeValue no pueden estar vacíos"})
		return
	}

	if newProduct.IsPublished != true && newProduct.IsPublished != false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "is_published debe ser un booleano válido"})
		return
	}

	_, err := time.Parse("01/01/2023", newProduct.Expiration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha de vencimiento incorrecto. Debe ser XX/XX/XXXX"})
		return
	}

	for _, p := range *products {
		if p.CodeValue == newProduct.CodeValue {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El valor de code_value ingresado ya existe"})
			return
		}
	}
	newProduct.ID = len(*products) + 1
	*products = append(*products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}
