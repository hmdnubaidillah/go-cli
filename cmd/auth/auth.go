package auth

import (
	pkg "cli/pkg"
	"encoding/base64"
	"fmt"
)

func HandleLogin() bool {

	for {
		fmt.Println()
		fmt.Println("=========Login=========")
		var username string
		var password string

		fmt.Print("Enter username : ")
		fmt.Scan(&username)

		fmt.Print("Enter password : ")
		fmt.Scan(&password)

		if user, isUserExist := pkg.UserCredentials[username]; isUserExist {

			encodedPassword, err := base64.StdEncoding.DecodeString(user.Password)

			if err == nil {

				if password == string(encodedPassword) {
					fmt.Println("Login success!")
					return true
				}

				fmt.Println("Password incorrect")
				return false
			}

			fmt.Println("Error", err.Error())
		}
		fmt.Println("Username not found")
	}
}

func HandleRegister() {

	for {
		fmt.Println()
		fmt.Println("=========Register=========")
		var username string
		var password string

		fmt.Print("Enter username :")
		fmt.Scan(&username)

		fmt.Print("Enter password :")
		fmt.Scan(&password)

		if _, isUserExist := pkg.UserCredentials[username]; isUserExist {
			fmt.Printf("username %s already exist\n", username)

		} else {
			// encode password
			encodedPass := base64.StdEncoding.EncodeToString([]byte(password))
			pkg.UserCredentials[username] = pkg.User{Username: username, Password: encodedPass}

			fmt.Println()
			fmt.Println("User created!")
			fmt.Println()
			return
		}
	}
}
