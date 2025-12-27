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
		HandleError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	artists, err := FetchArtists()
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Failed to fetch artists")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Failed to load template")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, artists); err != nil {
		HandleError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	buf.WriteTo(w)
}
