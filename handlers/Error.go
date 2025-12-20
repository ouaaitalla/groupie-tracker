package zone

import (
	"html/template"
	"net/http"
)

func HandleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	tmpl, _ := template.ParseFiles("templates/err.html")

	data := struct {
		Message string
		Status  int
	}{
		Message: message,
		Status:  status,
	}

	tmpl.Execute(w, data)
}
