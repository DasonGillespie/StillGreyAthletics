package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes
	http.HandleFunc("/", homeHandler)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	pgdata := PageData{
		Title: "Home",
	}

	renderTemplate(w, "home", pgdata)
}

type PageData struct {
	Title string
}

// Function to render HTML templates
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	// Parse the template files
	t, err := template.ParseFiles("templates/base.html", "static/"+tmpl+".html")
	if err != nil {
		http.Error(w, "Error loading template files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
