package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	app.InfoLog.Println("home called")
	files :=  []string{
					"./ui/html/home.page.tmpl",
					"./ui/html/base.layout.tmpl",
					"./ui/html/footer-partial.tmpl",
	}

	ts, err := template.ParseFiles(files ...)
	if err!=nil {

		app.ErrorLog.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err!=nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, "internal server error", 500)
	}


	//w.Write([]byte("hello from Snippetbox"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err:= strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil || id<1{
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "showing snippet for id..%d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not method", 405)
		return
	}

	w.Write([]byte("creating new snippet"))

}