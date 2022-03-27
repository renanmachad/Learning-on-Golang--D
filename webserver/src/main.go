package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	/*
	*	If the path of request is not "/hello"
	*	and the method is not "GET"
	* 	returns a 404 error respons to the user
	 */
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "hello!")
}

func formsHandler(w http.ResponseWriter, r *http.Request) {	
	if err:= r.ParseForm()
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formsHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
