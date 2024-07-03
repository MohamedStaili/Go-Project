package main

import (
	"fmt"
	"net/http"

	Config "github.com/MohamedStaili/Go-Project.git"
	"github.com/MohamedStaili/Go-Project.git/internal/handlers"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contactez-moi")
}

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	Config.Config()
}
