package main

import (
	"fmt"
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
	err := tpl.ExecuteTemplate(w, "index.html", Result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Result = ""
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		input := r.FormValue("text")
		banner := r.FormValue("banner")

		if len(input) >= 1000 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		result := fs.FinalPrint(input, banner)

		Result = result

		if len(Result) == 0 || Result == "incorrect banner" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if Result == "error in the file" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	var err error
	tpl, err = tpl.ParseGlob("html/*.html")
	if err != nil {
		panic(err)
	}
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiHandler)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
