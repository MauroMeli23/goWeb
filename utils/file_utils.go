package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var Products []Product

func LoadProducts() ([]Product, error) {
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo json:", err)
		return []Product{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Products); err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return []Product{}, err
	}

	return Products, nil
}
