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


type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Locationf struct {
    Indexf []struct {
		Id int `json:"id"`
        Locations []string `json:"locations"`
		Dates string `json:"dates"`
    } `json:"index"`
}

type Index struct {
	Locations []Location `json:"locations"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Loc struct {
	Locations []Locations `json:"index"`
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Data struct {
	Artist    Artists
	Locations Location
	Dates     Dates
	Relation  Relations
}



var ArtistsData []Artists
var LocationsData Loc
