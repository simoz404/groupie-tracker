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
	members, _ := strconv.Atoi(r.FormValue("Members"))

    for i := range utils.ArtistsData {
        utils.ArtistsData[i].FirstAlbum = utils.ArtistsData[i].FirstAlbum[len(utils.ArtistsData[i].FirstAlbum)-4:]
    }
	err = tmpl.Execute(w, struct {
		Artists []utils.Artists
		MinCD   int
		MaxCD   int
		MinFA   string
		MaxFA   string
		Members int
	}{
		Artists: utils.ArtistsData,
		MinCD:   minCD,
		MaxCD:   maxCD,
		MinFA:   minFA,
		MaxFA:   maxFA,
		Members: members,
	})
	if err != nil {
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
