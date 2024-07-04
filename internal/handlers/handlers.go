package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/MohamedStaili/Go-Project.git/config"
	model "github.com/MohamedStaili/Go-Project.git/internal/handlers/Model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	name := make(map[string]string)
	name["Root"] = "Mohamed"
	randerTemplate(w, "home", &model.TemplateData{StringData: name})
}

func Contact(w http.ResponseWriter, r *http.Request) {
	age := make(map[string]int)
	age["Root"] = 21
	randerTemplate(w, "contact", &model.TemplateData{IntData: age})
}

var appConfig *config.Config

func CreateTemplate(app *config.Config) {
	appConfig = app
}
func randerTemplate(w http.ResponseWriter, tmplName string, td *model.TemplateData) {
	templateCache := appConfig.TemplateCache
	tmpl, ok := templateCache[tmplName+".page.tmpl"]
	if !ok {
		http.Error(w, "template not exist", http.StatusInternalServerError)
		return
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
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
