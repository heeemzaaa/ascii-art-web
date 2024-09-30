package main

import (
	"html/template"
	"log"
	"net/http"

	fs "fs/ascii"
)

var (
	tpl    *template.Template
	Result string
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	Result = ""
	err := tpl.ExecuteTemplate(w, "index.html", Result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		input := r.FormValue("text")
		banner := r.FormValue("banner")

		result := fs.FinalPrint(input, banner)

		Result = result
		err := tpl.ExecuteTemplate(w, "index.html", Result)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
		}

	} else {
		http.Error(w, "Bad Request: Missing input", http.StatusBadRequest)
		return
	}
}

func main() {
	var err error
	tpl, err = tpl.ParseGlob("html/*.html")
	if err != nil {
		panic(err)
	}
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
