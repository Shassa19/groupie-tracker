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
	body, err := io.ReadAll(response.Body)
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
		members := strings.Join(artist.Members, ", ")
		fmt.Printf("Nom: %s,\n ID: %d,\n Image: %s,\n Membres: %v,\n Date de création: %d,\n Premier album: %s\n",
			artist.Name, artist.ID, artist.Image, members, artist.CreationDate, artist.FirstAlbum)
		fmt.Println("\n-----------------------------------------------------------------------------------")
	}
}
