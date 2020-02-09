package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page struct
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handleHTTPRequst(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/"):]
	// p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	fmt.Fprintf(w, "<html><body><h1>hello world</h1><p>path: %s</p></body></html>", r.URL.Path)
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/", handleHTTPRequst)
	log.Fatal(http.ListenAndServe(":4001", nil))
}
