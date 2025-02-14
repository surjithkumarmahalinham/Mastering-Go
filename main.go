package main

import (
	"fmt"
	"log"
	"net/http"
)

func formValidate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(w, "Parseform errror: %v", err)
	}
	fmt.Println(w, "post request success")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")


}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not support", http.StatusNotFound)
		return
	}
	fmt.Println(w, "hello!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.Handle("/form", http.HandlerFunc(formValidate))
	http.Handle("/hello", http.HandlerFunc(formValidate))

	fmt.Println("Starting server at port 8008")

	if err := http.ListenAndServe(":8008", nil); err != nil {
		log.Fatal(err)
	}
}
