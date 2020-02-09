package main

import (
	"io/ioutil"
	"path"
)

var pathToPages = path.Join("http-server", "pages")

// Page ...
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(path.Join(pathToPages, filename), p.Body, 0600)
}

// LoadPage ...
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(path.Join(pathToPages, filename))
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
