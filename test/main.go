package main

import (
	"fmt"
	"net/http"
)

type data struct {
	Text   string
	Banner string
}

var (
	d      data
	text   string
	banner string
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "file.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/process", process)
	http.HandleFunc("/show", showData)
	http.ListenAndServe(":8080", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	text = r.FormValue("txt")
	banner = r.FormValue("banner")
	d = data{Text: text, Banner: banner}

	http.Redirect(w, r, "/show", http.StatusSeeOther)
}

func CollectData() data {
	return d
}

func showData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Text: %s\nBanner: %s", d.Text, d.Banner)
}
