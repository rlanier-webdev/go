package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte //byte slice
}

type Homepage struct {
	Title string
	Pages []*Page // Slice of pointers to Page structs
}

var (
	pageDir = "./pages/"
	templates = template.Must(template.ParseFiles("edit.html", "view.html", "home.html"))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") // parse and compile the regular expression, and return a regexp
)

// This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error.
func (p *Page) save() error {
	filename := pageDir + p.Title + ".txt" // save the Page's Body to a text file.
	err := os.WriteFile(filename, p.Body, 0600)
	if err != nil {
		log.Printf("Error saving page %s: %v\n", p.Title, err)
		return err
	}
	return nil
}

// load the page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := os.ReadFile(pageDir + filename)

	if err != nil {
		log.Printf("Error reading file %s: %v\n", filename, err)

		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}



// render html files
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		log.Printf("Error rendering template %s: %v\n", tmpl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// allow users to VIEW a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	// if the requested Page doesn't exist, it should redirect the client to the edit Page so the content may be created
	if err != nil {
		log.Printf("Error handling view request for %s: %v\n", title, err)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

// allow users to EDIT a wiki page
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

// save page edits
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	newTitle := r.FormValue("title") // Get the new title from the form data
	p := &Page{Title: newTitle, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if newTitle != "" {
		http.Redirect(w, r, "/home/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}

}

func getAvailablePageTitles() ([]*Page, error) {
	var pages []*Page

	files, err := os.ReadDir(pageDir)
	if err != nil {
		log.Printf("Error reading directory %s: %v\n", pageDir, err)
		return nil, err
	}

	for _, file := range files {
		base := filepath.Base(file.Name())
		title := strings.TrimSuffix(base, ".txt")
		p, err := loadPage(title)
		if err != nil {
			return nil, err // Handle error loading individual pages
		}
		pages = append(pages, p)
	}
	return pages, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pages, err := getAvailablePageTitles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	homepage := &Homepage{Title: "Homepage", Pages: pages}
	renderTemplate(w, "home", homepage)
}

func main() {
	// handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
