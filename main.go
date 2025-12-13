package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
type Locations struct {
	Locations []string `json:"locations"`
}

var allArtists []Artist

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	return artists, err
}

func FetchLocation(id int) ([]string, error) {
	relationsURL := "https://groupietrackers.herokuapp.com/api/locations/"
	resp, err := http.Get(relationsURL + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rel Locations
	err = json.Unmarshal(body, &rel)
	if err != nil {
		return nil, err
	}

	return rel.Locations, nil
}

func handlerArtist(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var artist Artist
	for _, a := range allArtists {
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

	tmpl, err := template.ParseFiles("artist.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, allArtists)
}

func main() {
	var err error
	allArtists, err = FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/artist", handlerArtist)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
