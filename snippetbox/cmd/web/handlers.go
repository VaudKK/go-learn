package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"net/http"
	"strconv"

	"github.com/VaudKK/go-learn/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// this is now handled by Pat
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	// Pat doesnt strip the colon from the named capture key hence the full colon
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err == models.ErrorNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request){
	app.render(w,r,"create.page.tmpl",nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", "POST")
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	r.Body = http.MaxBytesReader(w,r.Body,4096)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w,http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errors := make(map[string]string)

	if strings.TrimSpace(title) == ""{
		errors["title"] = "This field cannot be blank"
	}else if utf8.RuneCountInString(title) > 100 { // utf8.RuneCount counts the number of chars and not bytes
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == ""{
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errors["expire"] = "This field cannot be blank"
	}else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	if len(errors) > 0 {
		fmt.Fprint(w,errors)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
