package main

import "fmt"

// we are store and print multiple number of data
func main() {

	func() {
		fmt.Println("Annonymous Function")
	}()
	showvalue("asfd", "adfd", "dssd")

	add := func(num1 int, num2 int) int {
		return num1 + num2
	}

}

func showvalue(s ...string) {
	fmt.Println(s)
}
