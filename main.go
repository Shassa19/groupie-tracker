package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

func main() {

	url := "https://groupietrackers.herokuapp.com/api/artists"

	// Detecte l'erreur lors de la requête
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer response.Body.Close()

	// Detecte lorsque ezrreur pdt la lecrure de la réponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return
	}

	var artists []Artist

	// Detecte reeur lors du décodage
	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return
	}

	// On affiche les donées recuperées
	for _, artist := range artists {
		fmt.Printf("Nom: %s, ID: %d, Image: %s, Membres: %v, Date de création: %d, Premier album: %s\n",
			artist.Name, artist.ID, artist.Image, artist.Members, artist.CreationDate, artist.FirstAlbum)
	}
}
