package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.adamuonu/internal/models"
)

// Define a home handler function which wrties a byte slice containing
// "Hello from snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.NotFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n\n", snippet)
	}
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// }

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	//log.Print(err.Error())
	// 	// app.errorLog.Print(err.Error())
	// 	// http.Error(w, "Internal Server Error", 500)
	// 	app.serverError(w, err)
	// 	return
	// }

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.Error(w, "Not found", http.StatusNotFound)
		app.NotFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	//log.Print(err.Error())
	// w.Write([]byte("Display a specific snippet..."))
	// fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
