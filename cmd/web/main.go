package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	Go_Projet.config()
}
