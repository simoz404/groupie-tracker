package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/handlers"
	"groupie-tracker/utils"
)

func main() {
	port := ":8080"

	http.HandleFunc("/assets/css/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.HandleError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		suffix := r.URL.Path[len("/assets/css/"):]

		if suffix != "index.css" && suffix != "artist.css" && suffix != "error.css" {
			utils.HandleError(w, "Access Forbidden", http.StatusForbidden)
			return
		}
		http.ServeFile(w, r, "./templates/css/"+suffix)
	})

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/filter", handlers.Filter)
	http.HandleFunc("/artist/", handlers.ArtistHandler)

	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}
