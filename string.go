package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "one"
	msg2 := "s"
	fmt.Println(strings.Compare(msg, msg2))

	contain := "golang programming language"
	fmt.Println(strings.Contains(contain, "golang"))

	lower := "msksurjith"
	fmt.Println(strings.ToUpper(lower))

	upper := "MSKSURJITH"
	fmt.Println(strings.ToLower(upper))

	fmt.Println(len(msg))
}
