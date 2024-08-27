package utils

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDate  string   `json:"concertDate"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Locations []string `json:"locations"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Data struct {
	Artist    Artists
	Locations Locations
	Dates     Dates
	Relation  Relations
}

var ArtistsData []Artists
