package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body  []byte
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(res, p)

	// fmt.Fprintf(res, "<h1>%s</h1><content>%s</content>", p.Title, p.Body)
	// fmt.Println(req.URL.Path)
}

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func editHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(res, p)
	// fmt.Fprintf(w,
	// 	"<h1>Editing %s</h1>" +
	// 	"<form action=\"/save/%s\" method=\"POST\">" +
	// 	"<textarea name=\"body\">%s</textarea></br>" +
	// 	"</form>",
	// 	, p.Title, p.Body)
}

func main() {
	// p1 := &Page{Title: "Test Page", Body: []byte("This is a simple Page.")}
	// p1.save()
	// p2, _ := loadPage("Test Page")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)

	http.ListenAndServe(":8080", nil)
}
