package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func main() {

	url := "https://groupietrackers.herokuapp.com/api"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return
	}

	var data ApiResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return
	}

	fmt.Printf("Données récupérées : %+v\n", data)
}
