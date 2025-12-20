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
}

var AllArtists []Artist

func HandlerArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, http.StatusBadRequest, "bad request")
		return
	}

	idStr := r.FormValue("id")

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
		HandleError(w, http.StatusNotFound, "artist not found")
		return
	}

	locations, _ := FetchLocation(artist.ID)
	for i, loc := range locations {
		locations[i] = FormatLocation(loc)
	}
	dates := FetchDate(artist.ID)
	dates = FormatDate(dates)
	rawRelation := FetchRelations(artist.ID)
	Relation := FormatRelations(rawRelation)
	data := struct {
		Artist    Artist
		Relations map[string][]string
		Locations []string
		Dates     []string
	}{
		Artist:    artist,
		Relations: Relation,
		Locations: locations,
		Dates:     dates,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
