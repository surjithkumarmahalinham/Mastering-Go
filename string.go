package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "one"
	msg2 := "s"

	fmt.Println(strings.Compare(msg, msg2))
	fmt.Println(len(msg))
}
