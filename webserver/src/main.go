package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	/*
	*	If the path of request is not "/hello"
	* 	returns a 404 error respons to the user
	 */
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
}

func formsHandler(w http.ResponseWriter, r *http.Request) {

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
