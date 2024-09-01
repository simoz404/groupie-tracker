package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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

	tmpl, err := template.ParseFiles("./templates/html/index.html")

	uniqueLocations := utils.LocactionsUnique()

	if err != nil {
		fmt.Println(err)
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, struct {
		Artists []utils.Artists
		LocationsUnique []string
	} {
		Artists: utils.ArtistsData,
		LocationsUnique: uniqueLocations,
	})
	if err != nil {
		utils.HandleError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
