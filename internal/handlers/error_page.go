package fs

import (
	"fmt"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	data := ErrorData{
		StatusCode: statusCode,
		Message:    message,
	}

	err := Tpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
