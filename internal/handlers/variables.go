package fs

import "text/template"

// global variables needed
var (
	Tpl *template.Template
	D   Data
)

type ErrorData struct {
	StatusCode int
	Message    string
}

type Data struct {
	Input  string
	Banner string
	Result string
}
