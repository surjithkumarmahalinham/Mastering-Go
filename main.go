package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path !=)
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.Handle("/form", formHandler)
	http.Handle("/hello", helloHandler)

	fmt.Println("Starting server at port 8008")

	if err := http.ListenAndServe(":8008", nil); err != nil {
		log.Fatal(err)
	}
}
