package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		utils.HandleError(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		utils.HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := utils.GetData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &utils.ArtistsData)
	if err != nil {
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/html/filter.html")
	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	minCD, _ := strconv.Atoi(r.FormValue("minCD"))
	maxCD, _ := strconv.Atoi(r.FormValue("maxCD"))
	minFA := r.FormValue("minFA")
	maxFA := r.FormValue("maxFA")
	m1, _ := strconv.Atoi(r.FormValue("m1"))
	m2, _ := strconv.Atoi(r.FormValue("m2"))
	m3, _ := strconv.Atoi(r.FormValue("m3"))
	m4, _ := strconv.Atoi(r.FormValue("m4"))
	m5, _ := strconv.Atoi(r.FormValue("m5"))
	m6, _ := strconv.Atoi(r.FormValue("m6"))
	m7, _ := strconv.Atoi(r.FormValue("m7"))
	m8, _ := strconv.Atoi(r.FormValue("m8"))
	location := r.FormValue("Locations")

	for i := range utils.ArtistsData {
		utils.ArtistsData[i].FirstAlbum = utils.ArtistsData[i].FirstAlbum[len(utils.ArtistsData[i].FirstAlbum)-4:]
	}

	utils.GetLocations(w)

	err = tmpl.Execute(w, struct {
		Artists  []utils.Artists
		Loc      utils.Loc
		MinCD    int
		MaxCD    int
		MinFA    string
		MaxFA    string
		M1       int
		M2       int
		M3       int
		M4       int
		M5       int
		M6       int
		M7       int
		M8       int
		Location string
		Is       bool
	}{
		Artists:  utils.ArtistsData,
		Loc:      utils.LocationsData,
		MinCD:    minCD,
		MaxCD:    maxCD,
		MinFA:    minFA,
		MaxFA:    maxFA,
		M1:       m1,
		M2:       m2,
		M3:       m3,
		M4:       m4,
		M5:       m5,
		M6:       m6,
		M7:       m7,
		M8:       m8,
		Location: location,
		Is:       false,
	})
	fmt.Println(utils.LocationsData.Locations)
	fmt.Println(err)
	if err != nil {
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
