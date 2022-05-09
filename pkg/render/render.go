package render

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/jerevick83/hotel_bookings/pkg/config"
	"github.com/jerevick83/hotel_bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func NewTemplate(a *config.AppConfig) {
	app = a
}
func TemplateRender(w http.ResponseWriter, gohtml string, td *models.TemplateData) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = BaseLayoutCache()
	}

	templates, ok := templateCache[gohtml]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	//makes a new pointer to the Buffer package
	buf := new(bytes.Buffer)

	//AddDefaultData adds data that should be rendered on every page
	td = AddDefaultData(td)
	//WriteTo ResponseWriter the data to be sent to the templates
	_ = templates.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

//BaseLayoutCache creates a template cache of the template pages in the templates root folder as a map
func BaseLayoutCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(sprig.FuncMap()).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		tempMatches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(tempMatches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}
	return myCache, nil
}
