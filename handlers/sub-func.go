package zone

import "strings"

func FormatLocation(s string) string {
	replacer := strings.NewReplacer(
		"_", " ",
		"-", " ",
	)
	return replacer.Replace(s)
}

func FormatRelations(rel map[string][]string) map[string][]string {
	formatted := make(map[string][]string)

	for place, dates := range rel {
		newPlace := FormatLocation(place)
		formatted[newPlace] = dates
	}

	return formatted
}

func FormatDate(dates []string) []string {
	for i, d := range dates {
		dates[i] = strings.TrimPrefix(d, "*")
	}
	return dates
}
