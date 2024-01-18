package manageproduct

import (
	helper "cli/helper"

	"fmt"
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
	helper.Scan()

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
	helper.Scan()

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
	helper.Scan()

	fmt.Println()
	fmt.Println("=========Edit Product==========")
	SeeProducts()
	fmt.Println("select product that'll be edited")

	fmt.Print("Enter number : ")
	productNumber, err := helper.ReadInt()

	if err == nil {
		for i := 0; i < len(Products); i++ {

			if productNumber-1 == i {

				fmt.Println()
				fmt.Println("1. Edit name")
				fmt.Println("2. Edit category")
				fmt.Println("3. Edit stock")
				fmt.Println("4. Edit price")
				fmt.Println()
				fmt.Println("Enter number : ")
				editChoice, err := helper.ReadInt()

				if err == nil {

					switch editChoice {
					case 1:
						editName(Products[i].Name, i)
					case 2:
						editCategory(Products[i].Category, i)
					case 3:
						editStock(Products[i].Stock, i)
					case 4:
						editPrice(Products[i].Price, i)
					default:
						fmt.Println("Please enter a valid input")
					}
				}
			}
		}
	}
}

func editName(name string, index int) {
	helper.Scan()

	fmt.Print("Enter new name : ")
	newName := helper.ReadString()
	Products[index].Name = newName
}

func editCategory(category string, index int) {
	helper.Scan()

	fmt.Print("Enter new category : ")
	newCategory := helper.ReadString()
	Products[index].Category = newCategory

}

func editStock(stock int, index int) {
	helper.Scan()

	fmt.Print("Enter new stock : ")
	newStock, err := helper.ReadInt()

	if err == nil {
		Products[index].Stock = newStock

	} else {
		fmt.Println("Please enter a valid input", err.Error())
	}

}

func editPrice(price int, index int) {
	helper.Scan()

	fmt.Print("Enter new price : ")
	newPrice, err := helper.ReadInt()

	if err == nil {
		Products[index].Price = newPrice

	} else {
		fmt.Println("Please enter a valid input", err.Error())
	}

}
