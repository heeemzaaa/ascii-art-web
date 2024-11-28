package fs

import (
	"net/http"

	fs "fs/internal/ascii"
)

// this function handles the request from the client and return a response
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		input := r.FormValue("text")
		banner := r.FormValue("banner")

		if len(input) >= 1000 {
			ErrorPage(w, http.StatusBadRequest, "Bad request !")
			return
		}

		result := fs.FinalPrint(input, banner)

		Result = result

		if len(Result) == 0 || Result == "incorrect banner" {
			ErrorPage(w, http.StatusBadRequest, "Bad request !")
			return
		}

		if Result == "error in the file" {
			ErrorPage(w, http.StatusInternalServerError, "Internal server error !")
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		ErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed !")
		return
	}
}
