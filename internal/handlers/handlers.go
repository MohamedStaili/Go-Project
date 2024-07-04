package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	randerTemplate(w, "home")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	randerTemplate(w, "contact")
}
func randerTemplate(w http.ResponseWriter, tmplName string) {

	tmpl, ok := templateCache[tmplName+".page.tmpl"]
	if !ok {
		http.Error(w, "template not exist", http.StatusInternalServerError)
		return
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, nil)
	buffer.WriteTo(w)
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))
		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
