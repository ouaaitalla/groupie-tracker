package zone

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchRelations(id int) map[string][]string {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Err in get:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Err in read:", err)
		return nil
	}

	var data Relations
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Err in json:", err)
		fmt.Println("Body:", string(body))
		return nil
	}

	return data.DatesLocations
}
