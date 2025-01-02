package main

import "fmt"

func main() {
	password := "12345654324"

	if len(password) > 10 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Invalid Password")
	}
}
