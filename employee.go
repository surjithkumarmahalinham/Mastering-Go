package main

func employee(name *string, age int) {
	if age > 65 {

	}
}

func main() {
	empName := "msk"
	age := 78
	employee(&empName, age)
}
