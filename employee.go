package main

func employee(name *string, age int) {
	if age > 65 {
		panic("Age cannot be greater then retirement age")
	}
}

func main() {
	empName := "msk"
	age := 78
	employee(&empName, age)
}
