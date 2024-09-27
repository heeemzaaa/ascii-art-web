package main

import (
	"html/template"
	"net/http"

	fs "fs/ascii"
)

var (
	tpl    *template.Template
	Result string
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Result = ""
	tpl.ExecuteTemplate(w, "index.html", Result)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		input := r.FormValue("text")
		banner := r.FormValue("banner")

		result := fs.FinalPrint(input, banner)

		Result = result
		tpl.ExecuteTemplate(w, "index.html", Result)
	}
}

func main() {
	var err error
	tpl, err = tpl.ParseGlob("html/*.html")
	if err != nil {
		panic(err)
	}
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/art", AsciiHandler)
	http.ListenAndServe(":5050", nil)
}
