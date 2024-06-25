package render

import (
	"fmt"
	"github.com/skarwa4491/bookings/models"
	"github.com/skarwa4491/bookings/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// var tc = make(map[string]*template.Template)

func RenderTemplateTest(response http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(response, nil)
	if err != nil {
		fmt.Println("error passing template", err.Error())
	}
}

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(response http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// create a template cache

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	td = AddDefaultTemplateData(td)
	err := t.Execute(response, td)
	if err != nil {
		log.Fatal("unable to parse template ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
