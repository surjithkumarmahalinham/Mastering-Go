package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "one"
	msg2 := "s"

	contain := "golang programming language"
	fmt.Println(strings.Contains(contain, "golang"))

	fmt.Println(strings.Compare(msg, msg2))

	fmt.Println(len(msg))
}
