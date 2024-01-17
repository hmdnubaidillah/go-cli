package main

import (
	auth "cli/auth"
	mainmenu "cli/main_menu"
	manageproduct "cli/manage_product"
	"encoding/base64"
	"fmt"
)

type Person struct {
	Id   int
	Name string
}

func setInitialCredentials() {
	auth.UserCredentials["kevin"] = auth.User{Username: "kevin", Password: base64.StdEncoding.EncodeToString([]byte("kevin123"))}
	auth.UserCredentials["john"] = auth.User{Username: "john", Password: base64.StdEncoding.EncodeToString([]byte("john123"))}
}

func setInitialProducts() {

	ID := len(manageproduct.Products)

	ID++
	newProduct1 := manageproduct.Product{
		Id:       ID,
		Name:     "Mouse Apik",
		Category: "Mouse",
		Stock:    10,
		Price:    120_000,
	}

	ID++
	newProduct2 := manageproduct.Product{
		Id:       ID,
		Name:     "Keyboard Apik",
		Category: "Keyboard",
		Stock:    12,
		Price:    160_000,
	}

	ID++
	newProduct3 := manageproduct.Product{
		Id:       ID,
		Name:     "Monitor Apik",
		Category: "Monitor",
		Stock:    5,
		Price:    1_000_000,
	}

	manageproduct.Products = append(manageproduct.Products, newProduct1)
	manageproduct.Products = append(manageproduct.Products, newProduct2)
	manageproduct.Products = append(manageproduct.Products, newProduct3)
}

func main() {
	// ID := 0

	// persons := []Person{}

	// ID++
	// person1 := Person{Id: ID, Name: "Kevin"}

	// ID++
	// person2 := Person{Id: ID, Name: "Johan"}

	// ID++
	// person3 := Person{Id: ID, Name: "Ahmed"}

	// persons = append(persons, person1)
	// persons = append(persons, person2)
	// persons = append(persons, person3)

	// // for i, p := range persons {
	// // 	fmt.Println(i, p)
	// // 	fmt.Println(p.Name == "Kevin")
	// // }

	// fmt.Println(persons[len(persons)-1])

	//==========================================
	setInitialCredentials()
	setInitialProducts()

	for {
		fmt.Println("==========Product Management==========")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Println()
		fmt.Print("Enter Command : ")

		var command int
		fmt.Scan(&command)

		switch command {

		case 1:

			if isLogin := auth.HandleLogin(); isLogin {
				mainmenu.HandleShowMainMenu()
			}

		case 2:
			auth.HandleRegister()

		case 0:
			fmt.Println("Goodbye!")
			return
		}
	}
}
