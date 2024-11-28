package fs

import (
	"net/http"
	"os"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my-css" {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}

	cssBytes, err := os.ReadFile("../static/css/styles.css")
	if err != nil {
		http.Error(w, "Error reading CSS file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/css")

	w.Write(cssBytes)
}
