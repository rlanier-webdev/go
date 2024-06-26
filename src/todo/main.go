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
	Body  []byte //byte slice
}

type Homepage struct {
	Title string
	Pages []*Page // Slice of pointers to Page structs
}

var (
	pageDir   = "./pages/"
	templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/home.html"))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9%-]+(?:%20[a-zA-Z0-9%-]+)*)$") // Updated pattern to allow spaces (%20) in the title
)

// This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error.
func (p *Page) save(title string) error {
	filename := pageDir + title + ".txt" // save the Page's Body to a text file.
	err := os.WriteFile(filename, p.Body, 0600)
	if err != nil {
		log.Printf("Error saving page %s: %v\n", p.Title, err)
		return err
	}
	return nil
}

// load the main page
func loadPage(title string) (*Page, error) {
	// Decode the title if necessary
	decodedTitle, err := url.PathUnescape(title)
	if err != nil {
		return nil, err
	}

	filename := decodedTitle + ".txt"

	body, err := os.ReadFile(pageDir + filename)

	if err != nil {
		log.Printf("Error reading file %s: %v\n", filename, err)

		return nil, err
	}

	return &Page{Title: decodedTitle, Body: body}, nil
}

// render html files
func renderTemplate(w http.ResponseWriter, templates string, data interface{}) {
	err := templates.ExecuteTemplate(w, templates+".html", data)
	if err != nil {
		log.Printf("Error rendering template %s: %v\n", templates, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Make Handler", r.URL.Path)
		// Here we will extract the page title from the Request, and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// allow users to VIEW a task page
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// Replace hyphens with spaces in the title for rendering on the page
	renderedTitle := strings.ReplaceAll(title, "-", " ")

	p, err := loadPage(renderedTitle)

	// if the requested Page doesn't exist, it should redirect the client to the edit Page so the content may be created
	if err != nil {
		log.Printf("Error handling view request for %s: %v\n", title, err)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// Pass the rendered title to the template
	p.Title = title

	renderTemplate(w, "view", p)
}

// allow users to EDIT a task page
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

	// Replace spaces with hyphens in the title
	newTitleWithHyphens := strings.ReplaceAll(newTitle, " ", "-")

	// Load the existing page
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// Check if the title has been changed
	if newTitleWithHyphens != title {
		// Delete the old note
		filename := pageDir + title + ".txt"
		if err := delFile(filename); err != nil {
			http.Error(w, "Failed to delete old note", http.StatusInternalServerError)
			return
		}
	}

	// Update the page with the new title and body
	p.Title = newTitle
	p.Body = []byte(body)

	// Save the updated page
	err = p.save(newTitleWithHyphens) // Corrected to use newTitleWithHyphens
	if err != nil {
		http.Error(w, "Failed to save note", http.StatusInternalServerError)
		return
	}
	// Redirect to the view page
	http.Redirect(w, r, "/view/"+newTitleWithHyphens, http.StatusSeeOther)
}


func delHandler(w http.ResponseWriter, r *http.Request, title string) {
	filename := pageDir + title + ".txt"
	if err := delFile(filename); err != nil {
		log.Printf("Error deleting file %s: %v\n", title, err)
		if os.IsNotExist(err) {
			// File does not exist, redirect to edit page
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
			return
		}
		// Handle other errors appropriately
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Redirect to the home page after successful deletion
	http.Redirect(w, r, "/home/", http.StatusFound)
}

func delFile(filename string) error {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// File does not exist, return without error
		return nil
	} else if err != nil {
		// Other errors, return the error
		return err
	}

	// Attempt to remove the file
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
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/del/", makeHandler(delHandler))

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
