package render

import (
	"bytes"
	"github.com/shynggys9219/goBookingProject/config"
	"github.com/shynggys9219/goBookingProject/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	// rebuild page everytime
	if app.UseCache {
		//get the template cache from the app config in main (config.AppConfig)
		tc = app.TemplateCach
	} else {
		// rebuilding template cache
		tc, _ = CreateTemplateCache()
		//if err
	}

	// get requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCach := make(map[string]*template.Template)
	// another way to create a map
	myCache := map[string]*template.Template{}

	// get all of the files ending .page.tmpl from templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println(err)
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// return last element (filaname) of the path
		name := filepath.Base(page)
		// ts - template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		log.Println("Creating a cached page", page, ts)
		myCache[name] = ts
	}
	return myCache, nil
}
