package main

// Define interface
type Human interface {
	Username() string
	Profession() string
}

// Define struct
type User1 struct {
	Age    int
	Gender string
}

type User2 struct {
	Age    int
	Gender string
}

// Define Methods
func (j User1) Profession() string {
	return "Front-end Developer"
}

func (j User1) Username() string {
	return "Abc"
}

func (a User2) Profession() string {
	return "Backend Developer"
}

func (a User2) Username() string {
	return "Xyz"
}

// User Information

func main() {
	// Instance of struct
	John := User1{
		Age:    27,
		Gender: "Male",
	}

	PrintInfo(John)

	Lara := User2{
		Age:    28,
		Gender: "Female",
	}
	PrintInfo(Lara)
}
