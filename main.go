package main

import (
		"fmt"
		"app/models"
)

func main() {
	products, erro := models.GetAll()

	if erro != nil {
		fmt.Println("Ops! Ocorreu um erro")
	}

	fmt.Println("Lista de produtos:\n")

	fmt.Println(products)
}