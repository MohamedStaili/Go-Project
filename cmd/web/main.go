package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contactez-moi")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	Go_Projet.config() // Calling the config function from Go_Projet package
}
