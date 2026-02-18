package main

import (
	"fmt"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the about page"))
	fmt.Println("Welcome to the about page")
}
func main() {
	http.HandleFunc("/", aboutHandler)
	http.ListenAndServe(":1234", nil)
}
