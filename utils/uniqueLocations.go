package utils

import (
	"encoding/json"
	"net/http"
	"sort"
)

func LocactionsUnique() []string {
	var w http.ResponseWriter
	var Locations Locationf
	data, err := GetData("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = json.Unmarshal(data, &Locations)
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
	}
	uniqueLocations := make(map[string]bool)

	for _, item := range Locations.Indexf {
		for _, location := range item.Locations {
			uniqueLocations[location] = true
		}
	}

	var locationSlice []string
	for location := range uniqueLocations {
		locationSlice = append(locationSlice, location)
	}

	sort.Strings(locationSlice)
	return locationSlice
}
