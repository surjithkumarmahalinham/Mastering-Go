package main

import "fmt"

func main() {

	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	var peoples = [3]string{"msk", "vicky", "diva"}

	for item := 0; item < len(peoples); item++ {
		fmt.Println(item, peoples[item])
	}

	fmt.Println(numbers)
}
