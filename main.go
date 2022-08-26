package main

import (
		"fmt"
		"./models"
)

func main() {
	products, erro := models.GetAll()

	if erro != nil {
		fmt.Println("Ops! Ocorreu um erro")
	}

	fmt.Println("Lista de producots:\n")

	fmt.Println(products)
}