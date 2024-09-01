package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/utils"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/artist/"):]
	_, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(w, "Invalid Artist ID", http.StatusBadRequest)
		return
	}

	data, err := utils.GetData("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		utils.HandleError(w, "Failed to retrieve artist data", http.StatusInternalServerError)
		return
	}

	artist := utils.Artists{}
	err = json.Unmarshal(data, &artist)
	if err != nil {
		utils.HandleError(w, "Failed to unmarshal artist data", http.StatusInternalServerError)
		return
	}

	if artist.Id == 0 {
		utils.HandleError(w, "Artist not found", http.StatusNotFound)
		return
	}

	locationData, err := utils.GetData("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		utils.HandleError(w, "Failed to retrieve artist data", http.StatusInternalServerError)
		return
	}

	locations := utils.Location{}
	err = json.Unmarshal(locationData, &locations)
	if err != nil {
		utils.HandleError(w, "Failed to unmarshal artist data", http.StatusInternalServerError)
		return
	}

	datesData, err := utils.GetData("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		utils.HandleError(w, "Failed to retrieve artist data", http.StatusInternalServerError)
		return
	}

	dates := utils.Dates{}
	err = json.Unmarshal(datesData, &dates)
	if err != nil {
		utils.HandleError(w, "Failed to unmarshal artist data", http.StatusInternalServerError)
		return
	}

	relationData, err := utils.GetData("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		utils.HandleError(w, "Failed to retrieve artist data", http.StatusInternalServerError)
		return
	}

	relation := utils.Relations{}
	err = json.Unmarshal(relationData, &relation)
	if err != nil {
		utils.HandleError(w, "Failed to unmarshal artist data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/html/artist.html")
	if err != nil {
		utils.HandleError(w, "Failed to parse template", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	results := utils.Data{
		Artist:    artist,
		Locations: locations,
		Dates:     dates,
		Relation:  relation,
	}

	err = tmpl.Execute(w, results)
	if err != nil {
		utils.HandleError(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
