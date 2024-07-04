package main

import (
	"fmt"
	"net/http"

	"github.com/MohamedStaili/Go-Project.git/config"
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
	var appConfig config.Config
	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		panic((err))
	}
	appConfig.TemplateCache = templateCache
	appConfig.port = ":8080"
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)
	fmt.Println("Starting server on port", appConfig.port)
	err := http.ListenAndServe(appConfig.port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
