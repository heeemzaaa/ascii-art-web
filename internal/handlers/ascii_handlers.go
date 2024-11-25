package fs

import (
	"net/http"

	fs "fs/internal/ascii"
)

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
