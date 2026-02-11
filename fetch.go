package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

var (
	artists   []Artist
	relations []Relation
)

func fetchAPI(url string, target interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Fatal(err)
	}
}

func loadData() {
	fetchAPI(baseURL+"/artists", &artists)

	var rel struct {
		Index []Relation `json:"index"`
	}
	fetchAPI(baseURL+"/relation", &rel)
	relations = rel.Index
}
