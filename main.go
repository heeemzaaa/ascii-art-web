package main

import (
	"html/template"
	"net/http"

	fs "fs/ascii"
)

type Data struct {
	Input  string
	Banner string
	Result string
}

var (
	output Data
	tpl    *template.Template
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	output = Data{
		Input:  "",
		Banner: "",
		Result: "",
	}
	tpl.ExecuteTemplate(w, "index.html", output)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		input := r.FormValue("text")
		banner := r.FormValue("banner")

		result := fs.FinalPrint(input, banner)

		output = Data{
			Input:  input,
			Banner: banner,
			Result: result,
		}
		tpl.ExecuteTemplate(w, "index.html", output)
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
