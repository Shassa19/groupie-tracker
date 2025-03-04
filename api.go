package api

import (
	"encoding/json"
	"fmt"
	apiRelation "groupie-tracker/ApiRelation"
	"io"
	"net/http"
)

// On définit la structure de l'Artist
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relation     string   `json:"relations"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var relation apiRelation.Relation

// Fonction pour récupérer les données depuis l'API
func FetchArtists(url string) ([]Artist, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la requête : %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la lecture de la réponse : %v", err)
	}

	var fetchedArtists []Artist
	err = json.Unmarshal(body, &fetchedArtists)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors du décodage JSON : %v", err)
	}
	return fetchedArtists, nil

}
