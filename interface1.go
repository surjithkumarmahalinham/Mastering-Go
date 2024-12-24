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

}

func main() {

}
