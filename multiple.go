package main

import "fmt"

// we are store and print multiple number of data
func main() {
	showvalue("asfd", "adfd", "dssd")
}

func showvalue(s ...string) {
	fmt.Println(s)
}
