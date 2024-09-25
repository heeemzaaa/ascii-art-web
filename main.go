package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	fs "fs/ascii"
)

type PageVariables struct {
	Input  string
	Result string
}

var tpl *template.Template

func main() {
	var err error
	tpl, err = template.ParseGlob("*.html")
	if err != nil {
		panic(err)
	}

	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/art", processForm)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	vars := PageVariables{
		Input:  "",
		Result: "",
	}
	renderTemplate(w, vars)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		banner := r.FormValue("banner")
		input := r.FormValue("text")

		result := finalPrint(input, banner)

		vars := PageVariables{
			Input:  input,
			Result: result,
		}
		renderTemplate(w, vars)
	}
}

func renderTemplate(w http.ResponseWriter, vars PageVariables) {
	if err := tpl.ExecuteTemplate(w, "index.html", vars); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}

func finalPrint(text string, banner string) string {
	name := ""
	if banner == "thinkertoy" || banner == "standard" || banner == "shadow" {
		name = banner
	} else {
		fmt.Println("incorrect banner")
		return ""
	}
	file := fs.Read_file(name)
	if file == nil {
		return ""
	}
	line := text
	if !fs.Is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return ""
	}
	if len(line) < 1 {
		return ""
	}
	finalResult := ""
	lines_count := fs.Count_next_line(line)
	splitted_line := strings.Split(line, "\\n")
	splitted_line, lines_count = fs.Cleaned_split(splitted_line, lines_count)
	finalResult = fs.Print_art(file[1:], splitted_line, lines_count)
	return finalResult
}
