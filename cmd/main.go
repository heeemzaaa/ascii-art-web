package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	h "fs/internal/handlers"
)

func main() {
	var err error
	h.Tpl, err = template.ParseFiles("../template/index.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/my-css", h.CssHandler)
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/ascii-art", h.AsciiHandler)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
