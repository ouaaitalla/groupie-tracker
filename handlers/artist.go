package zone

import (
	"html/template"
	"net/http"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
}

var AllArtists []Artist

func HandlerArtist(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var artist Artist
	for _, a := range AllArtists {
		if a.ID == id {
			artist = a
			break
		}
	}

	locations, err := FetchLocation(artist.ID)
	if err != nil {
		locations = []string{"Locations not available"}
	}

	data := struct {
		Artist    Artist
		Locations []string
	}{
		Artist:    artist,
		Locations: locations,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
