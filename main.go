package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	fs "fs/ascii"
)

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

	filename := "result.txt"
	resultFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating a file:", err)
		return ""
	}
	defer resultFile.Close()

	_, err = resultFile.WriteString(finalResult)
	if err != nil {
		fmt.Println("Error writing the result")
		return ""
	}
	return finalResult
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "file.html")
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	text := r.FormValue("txt")
	banner := r.FormValue("banner")

	// Call your processing function
	result := finalPrint(text, banner)
	// Display the result on a new page
	fmt.Fprintf(w, "<h1>Result</h1>")
	fmt.Fprintf(w, "<p>%s</p>", result)
	fmt.Fprintf(w, "<a href='/'>Go back</a>")
}

func Print() string {
	return "OK"
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}
