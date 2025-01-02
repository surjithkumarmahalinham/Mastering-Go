package main

func employee(name *string, age int) {

}

func main() {
	empName := "msk"
	age := 78
	employee(&empName, age)
}
