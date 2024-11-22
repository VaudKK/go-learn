package main

import (
	"html/template"
	"path/filepath"

	"github.com/VaudKK/go-learn/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
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

		ts, err := template.ParseFiles(page)
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
