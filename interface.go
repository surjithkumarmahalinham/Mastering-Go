package main

import "fmt"

type Animal interface {
	GetInfo() string
}

func (c Cat) GetInfo() string {
	return fmt.Sprintf("Cat:%s, Color:%s", c.Name, c.Color)
}

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) GetInfo() string {
	return fmt.Sprintf("Dog: %s, Breed: %s", d.Name, d.Breed)
}
func printAnimalInfo(animal Animal) {
	fmt.Println(animal.GetInfo())
}

func main() {
	cat := Cat{Name: "Jack", Color: "Gray"}
	dog := Dog{Name: "Domy", Breed: "Labaror"}
	printAnimalInfo(cat)
	printAnimalInfo(dog)
}
