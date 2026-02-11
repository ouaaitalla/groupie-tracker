package main

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Code    int
	Message string
}

var tmpl = template.Must(template.ParseFiles("templates/error.html"))

func HandelError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	data := ErrorData{
		Code:    code,
		Message: message,
	}
	tmpl.Execute(w, data)
}
