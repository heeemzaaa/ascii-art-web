package fs

import "net/http"

// this function represents the home page after running the server
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := Tpl.ExecuteTemplate(w, "index.html", Result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Result = ""
}
