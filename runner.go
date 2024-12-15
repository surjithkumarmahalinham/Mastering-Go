package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
	var one, two, three, four, five int := 1, 2, 3, 4, 5
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	
	fmt.Println(five)
}
