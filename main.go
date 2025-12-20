package main

import (
	"fmt"
	"net/http"

	zone "zone/handlers"
)

func main() {
	var err error
	zone.AllArtists, err = zone.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}
	http.HandleFunc("/static/", zone.HandleStatic)
	http.HandleFunc("/", zone.HandlerHome)
	http.HandleFunc("/artist/", zone.HandlerArtist)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
