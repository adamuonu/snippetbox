package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function which wrties a byte slice containing
// "Hello from snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.NotFound(w)
		return
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//log.Print(err.Error())
		// app.errorLog.Print(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		//log.Print(err.Error())
		// app.errorLog.Print(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.Error(w, "Not found", http.StatusNotFound)
		app.NotFound(w)
		return
	}
	// w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
