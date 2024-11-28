package fs

import (
	"net/http"
)


func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	data := ErrorData{
		StatusCode: statusCode,
		Message:    message,
	}
	Tpl.ExecuteTemplate(w, "error.html", data)
}