package main

import (

	"fmt"
	"net/http"
	"strconv"
	"log"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files :=  []string{
					"./ui/html/home.page.tmpl",
					"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files ...)
	if err!=nil {

		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err!=nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
	}


	//w.Write([]byte("hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err:= strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil || id<1{
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "showing snippet for id..%d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not method", 405)
		return
	}

	w.Write([]byte("creating new snippet"))

}