package fs

import (
	"net/http"
)

type ErrorData struct {
	StatusCode int
	Message    string
}

func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	data := ErrorData{
		StatusCode: statusCode,
		Message:    message,
	}
	Tpl.ExecuteTemplate(w, "error.html", data)
}