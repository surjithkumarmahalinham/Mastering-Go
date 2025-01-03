// A struct is used to create a collection of members of different data types, into a single variable.

package main

type Person struct {
	name    string
	age     int
	address string
}

type Programmer struct {
	Person
	isProgrammer bool
}

func main() {
	h := Person{
		name:    "msk",
		age:     28,
		address: "chennai",
	}

}
