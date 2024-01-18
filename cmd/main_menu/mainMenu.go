package mainmenu

import (
	manageproduct "cli/cmd/manage_product"
	"fmt"
)

func HandleShowMainMenu() {
	for {
		fmt.Println()
		fmt.Println("=========Main Menu==========")
		fmt.Println("1. See Product")
		fmt.Println("2. Add Product")
		fmt.Println("3. Remove Product")
		fmt.Println("4. Edit Product")
		fmt.Println("0. Back")
		fmt.Println()

		var command int
		fmt.Print("Enter command : ")
		fmt.Scan(&command)

		switch command {
		case 1:
			manageproduct.SeeProducts()

		case 2:

			manageproduct.AddProduct()

		case 3:
			manageproduct.DeleteProduct()

		case 4:
			manageproduct.EditProduct()

		case 0:

			return
		default:
			fmt.Println("Please enter a valid command")
		}
	}

}
