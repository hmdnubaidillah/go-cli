package manageproduct

import (
	"bufio"
	"cli/helper"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Name, Category   string
	Id, Stock, Price int
}

var Products = []Product{}

func SeeProducts() {
	fmt.Println()
	fmt.Println("==========Products==========")

	for _, v := range Products {
		fmt.Printf("Id : %d\n", v.Id)
		fmt.Printf("Name : %s\n", v.Name)
		fmt.Printf("Category : %s\n", v.Category)
		fmt.Printf("Stock : %d\n", v.Stock)
		fmt.Printf("Price : %d\n", v.Price)
		fmt.Println()
	}
}

func AddProduct() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	temp := []string{}

	fmt.Println()
	fmt.Println("=========Add Product==========")

	var productAmount int
	var name string
	var category string

	fmt.Print("Enter amount of product that'll be added : ")
	fmt.Scan(&productAmount)

	isProductExist := false

	for _, v := range Products {
		if v.Name == name {
			isProductExist = true
		}
	}

	if !isProductExist {
		for i := 0; i < productAmount; i++ {

			fmt.Println("Name : ")
			name = helper.ReadString()

			fmt.Println("Category : ")

			category = helper.ReadString()

			fmt.Println("Stock : ")
			stock, errStock := helper.ReadInt()

			fmt.Println("Price : ")
			price, errPrice := helper.ReadInt()

			if errPrice == nil && errStock == nil {

				idOfLastIndexProducts := Products[len(Products)-1]

				idOfLastIndexProducts.Id++
				newProduct := Product{
					Id:       idOfLastIndexProducts.Id,
					Name:     name,
					Category: category,
					Stock:    stock,
					Price:    price,
				}
				Products = append(Products, newProduct)

			} else {
				fmt.Println("Error", errPrice.Error(), errStock.Error())
			}
		}
	} else {
		fmt.Println(name, "already exist")
	}

	names := append(temp, name)
	fmt.Printf("Products added : %v\n", names)

}

func DeleteProduct() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("==========Delete Product==========")
	SeeProducts()

	fmt.Println()
	fmt.Println("select product that'll be deleted")
	productNumber, err := helper.ReadInt()

	if err == nil {
		fmt.Println("Are you sure (y/n)")
		deleteConfirmation := helper.ReadString()

		if strings.ToLower(deleteConfirmation) == "y" {

			for i, v := range Products {

				if productNumber == i {
					fmt.Println(v)
				}
			}
		} else if strings.ToLower(deleteConfirmation) == "n" {
			return

		} else {
			fmt.Println("Please enter a valid command")
		}

	} else {

		fmt.Println("Error", err.Error())
	}

}

func EditProduct() {

}
