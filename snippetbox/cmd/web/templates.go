package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/VaudKK/go-learn/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	//initialize a new map to act as cache
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}

	//loop through the pages one by one
	for _, page := range pages {
		//extract the file name
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before
		// calling the ParseFiles() method.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		// use the parseGlob method to add any 'layout' templates to the
		// template set. In this case the 'base' layout at the moment
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))

		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
