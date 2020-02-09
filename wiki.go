package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

// Page struct
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(path.Join("pages", filename), p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(path.Join("pages", filename))
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func main() {
	http.HandleFunc("/edit/", handleEdit)
	http.HandleFunc("/save/", handleSave)
	http.HandleFunc("/view/", handleView)
	log.Fatal(http.ListenAndServe(":4001", nil))
}
