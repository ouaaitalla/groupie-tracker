package zone

import (
	"bytes"
	"html/template"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound, "Page not found")
		return
	}
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusBadRequest, "bad request")
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, AllArtists); err != nil {
		HandleError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	buf.WriteTo(w)
}
