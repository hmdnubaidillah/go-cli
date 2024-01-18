package main

import (
	auth "cli/cmd/auth"
	mainmenu "cli/cmd/main_menu"
	"cli/pkg"
	"encoding/base64"
	"fmt"
)

type Person struct {
	Id   int
	Name string
}

func setInitialCredentials() {
	pkg.UserCredentials["kevin"] = pkg.User{Username: "kevin", Password: base64.StdEncoding.EncodeToString([]byte("kevin123"))}
	pkg.UserCredentials["john"] = pkg.User{Username: "john", Password: base64.StdEncoding.EncodeToString([]byte("john123"))}
}

func setInitialProducts() {

	ID := len(pkg.Products)

	ID++
	newProduct1 := pkg.Product{
		Id:       ID,
		Name:     "Mouse Apik",
		Category: "Mouse",
		Stock:    10,
		Price:    120_000,
	}

	ID++
	newProduct2 := pkg.Product{
		Id:       ID,
		Name:     "Keyboard Apik",
		Category: "Keyboard",
		Stock:    12,
		Price:    160_000,
	}

	ID++
	newProduct3 := pkg.Product{
		Id:       ID,
		Name:     "Monitor Apik",
		Category: "Monitor",
		Stock:    5,
		Price:    1_000_000,
	}

	pkg.Products = append(pkg.Products, newProduct1)
	pkg.Products = append(pkg.Products, newProduct2)
	pkg.Products = append(pkg.Products, newProduct3)
}

func main() {
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
