package fs

import "net/http"

// this function represents the home page after running the server
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, 404, "Page not found !")
		return
	}
	err := Tpl.ExecuteTemplate(w, "index.html", Result)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "Internal server error !")
		return
	}

	Result = ""
}
