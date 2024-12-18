package main

import "fmt"

func main() {
	password := "12345654324"

	if len(password) > 10 {
		fmt.Println("Valid Password")
	} else if password == "" {
		fmt.Println("Password is Required")
	} else {
		fmt.Println("Invalid Password")
	}
}
