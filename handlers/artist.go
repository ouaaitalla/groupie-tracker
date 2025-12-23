package zone

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var AllArtists []Artist

func HandlerArtist(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	if idStr == "" {
		HandleError(w, http.StatusNotFound, "Page not found")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	var artist Artist
	found := false
	for _, a := range AllArtists {
		if a.ID == id {
			artist = a
			found = true
			break
		}
	}

	if !found {
		HandleError(w, http.StatusNotFound, "Artist not found")
		return
	}

	locations, _ := FetchLocation(artist.ID)
	dates := FetchDate(artist.ID)
	dates = FormatDate(dates)
	relations := FetchRelations(artist.ID)

	for i, loc := range locations {
		locations[i] = FormatLocation(loc)
	}

	data := struct {
		Artist    Artist
		Relations map[string][]string
		Locations []string
		Dates     []string
	}{
		Artist:    artist,
		Relations: relations,
		Locations: locations,
		Dates:     dates,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Failed to load template")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		HandleError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	buf.WriteTo(w)
}
