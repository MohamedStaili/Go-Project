package main

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	randerTemplate(w, "home")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	randerTemplate(w, "contact")
}
func randerTemplate(w http.ResponseWriter, tmpl string) {
	t, arr := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if arr != nil {
		http.Error(w, arr.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
