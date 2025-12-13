package zone

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Locations struct {
	Locations []string `json:"locations"`
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
