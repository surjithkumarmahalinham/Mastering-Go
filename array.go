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

	// slice

	num := []int{10, 20}

	fmt.Println(num)
	num = append(num, 30, 40, 50)
	fmt.Println(num)
	fmt.Println(numbers)

	// make

	num1 := make([]int, 5, 20)

	num1[0] = 10
	num1[1] = 20
	num1[2] = 30
	num1[3] = 40

	fmt.Println(num1)
	fmt.Println(len(num1))
	fmt.Println(cap(num1))

	userData := map[string]int{

		"vky": 100,
		"dva": 100,
	}

	for prop, value := range userData {
		fmt.Println(prop, value)
	}
}
