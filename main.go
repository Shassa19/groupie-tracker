package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// On définit la structure de l'Artist
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var artists []Artist

// Fonction pour récupérer les données depuis l'API
func fetchArtists(url string) ([]Artist, error) {
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

// Fonction pour afficher les artistes
func displayArtists(artists []Artist) {
	for _, artist := range artists {
		members := strings.Join(artist.Members, ", ")
		fmt.Printf("Nom: %s,\n ID: %d,\n Image: %s,\n Membres: %v,\n Date de création: %d,\n Premier album: %s\n",
			artist.Name, artist.ID, artist.Image, members, artist.CreationDate, artist.FirstAlbum)
		fmt.Println("\n-----------------------------------------------------------------------------------")
	}
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	var err error
	artists, err = fetchArtists(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	displayArtists(artists)
}
