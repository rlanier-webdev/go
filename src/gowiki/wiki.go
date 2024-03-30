package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte //byte slice
}

// This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error.
func (p *Page) save() error {
	filename := p.Title + ".txt" // save the Page's Body to a text file.
	return os.WriteFile(filename, p.Body, 0600)
}

// load the page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// render html files
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// allow users to VIEW a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	// if the requested Page doesn't exist, it should redirect the client to the edit Page so the content may be created
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

// allow users to EDIT a wiki page
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

// save page edits
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// handlers
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
