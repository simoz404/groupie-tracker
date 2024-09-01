package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetLocations(w http.ResponseWriter) {
	data, err := GetData("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &LocationsData)
	if err != nil {
		HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
