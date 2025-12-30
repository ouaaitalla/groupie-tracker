package zone

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func FetchDate(id int) ([]string, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Dates
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Dates, nil
}
