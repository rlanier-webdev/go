package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

type Homepage struct {
	Title string
	Pages []*Page
}

var (
	pageDir   = "./pages/"
	templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/home.html"))
	validPath = regexp.MustCompile("^/(edit|save|view|del)/([a-zA-Z0-9%-]+(?:%20[a-zA-Z0-9%-]+)*)$")
)

func (p *Page) save(title string) error {
	filename := filepath.Join(pageDir, title+".txt")
	err := os.WriteFile(filename, p.Body, 0600)
	if err != nil {
		log.Printf("Error saving page %s: %v\n", p.Title, err)
		return err
	}
	return nil
}

func loadPage(title string) (*Page, error) {
	decodedTitle, err := url.PathUnescape(title)
	if err != nil {
		return nil, err
	}

	filename := filepath.Join(pageDir, decodedTitle+".txt")
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file %s: %v\n", filename, err)
		return nil, err
	}

	return &Page{Title: decodedTitle, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		log.Printf("Error rendering template %s: %v\n", tmpl, err)
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

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderedTitle := strings.ReplaceAll(title, "-", " ")
	p, err := loadPage(renderedTitle)
	if err != nil {
		log.Printf("Error handling view request for %s: %v\n", title, err)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	p.Title = title
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderedTitle := strings.ReplaceAll(title, "-", " ")
	p, err := loadPage(renderedTitle)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	newTitle := r.FormValue("title")

	newTitleWithHyphens := strings.ReplaceAll(newTitle, " ", "-")
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	if newTitleWithHyphens != title {
		filename := filepath.Join(pageDir, title+".txt")
		if err := delFile(filename); err != nil {
			http.Error(w, "Failed to delete old note", http.StatusInternalServerError)
			return
		}
	}

	p.Title = newTitle
	p.Body = []byte(body)

	err = p.save(newTitleWithHyphens)
	if err != nil {
		http.Error(w, "Failed to save note", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func delHandler(w http.ResponseWriter, r *http.Request, title string) {
	filename := filepath.Join(pageDir, title+".txt")
	if err := delFile(filename); err != nil {
		log.Printf("Error deleting file %s: %v\n", title, err)
		if os.IsNotExist(err) {
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func delFile(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	if err := os.Remove(filename); err != nil {
		return err
	}

	return nil
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
			return nil, err
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
	http.HandleFunc("/", homeHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/del/", makeHandler(delHandler))

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
