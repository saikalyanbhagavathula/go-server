package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successful\n")

	name := r.FormValue("name")
	fmt.Fprintf(w, "The name value in the form is: %v\n", name)
	gender := r.FormValue("gender")
	fmt.Fprintf(w, "The gender value in the string is: %v\n", gender)
	return
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "You are in HELLO Page")
}
func main() {
	fileserver := http.FileServer(http.Dir("./html"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting Server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("err")
	}
}
