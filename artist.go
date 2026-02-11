package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandelError(w, 405, "Method Not Allowed")
		return
	}
	if r.URL.Path != "/" {
		HandelError(w, 404, "Page Not Found")
		return
	}
	err := templates.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		HandelError(w, 500, "internal server error")
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandelError(w,404, "page not found")
		return
	}

	var artist Artist
	var relation Relation

	for _, a := range artists {
		if a.ID == id {
			artist = a
			break
		}
	}

	for _, r := range relations {
		if r.ID == id {
			relation = r
			break
		}
	}

	data := struct {
		Artist   Artist
		Relation Relation
	}{
		Artist:   artist,
		Relation: relation,
	}

	err = templates.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
