package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type artistsInfo struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Memebers     []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstalbum"`
	Locations    string   `json:"locations"`
	ConsertDates string   `json:"consertdates"`
	Rolations    string   `json:"rolations"`
}

type BandInfo struct {
    Id             int                 `json:"id"`
	Image        string   			`json:"image"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/Infos", Infos)
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

func GetData(url string, method string) []byte {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	// defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	return body
}

func Home(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	data := GetData(url, "GET")
	s := []artistsInfo{}
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Fatal("ERROR ENCODING JSON")
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, s)
}

func Infos(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/relation" + "/" + r.FormValue("band")
	data := GetData(url, "GET")
	var bandInfo BandInfo
	err := json.Unmarshal(data, &bandInfo)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
	}
	bandInfo.Image = r.FormValue("bandimage")
	tmpl, _ := template.ParseFiles("bandInfos.html")
	tmpl.Execute(w, bandInfo)
}