package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"regexp"
)

// pre-parse templates
var templates = template.Must(template.ParseFiles(
	path.Join("http-server", "templates", "edit.html"),
	path.Join("http-server", "templates", "view.html"),
))

// only allow specific templates (otherwise insecure)
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func handleEdit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func handleSave(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func handleView(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func main() {
	http.HandleFunc("/edit/", makeHandler(handleEdit))
	http.HandleFunc("/save/", makeHandler(handleSave))
	http.HandleFunc("/view/", makeHandler(handleView))
	log.Fatal(http.ListenAndServe(":4001", nil))
}
