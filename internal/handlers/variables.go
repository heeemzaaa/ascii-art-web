package fs

import "text/template"

// global variables needed
var (
	Tpl    *template.Template
	Result string
)

type ErrorData struct {
	StatusCode int
	Message    string
}
