package apiRelation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// On définit la structure de l'Artist
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var relation Relation

// Fonction pour récupérer les données depuis l'API
func FetchInfos(id int) (Relation, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + string(id))

	// On récupère les relations depuis l'API relation via la fonction fetcInfos du package apiRelation
	var erro error
	if erro != nil {
		fmt.Println("Erreur lors de la récupération des relations :", erro)
	} else {
		fmt.Println("Relation récupérée :", relation)
	}

	if err != nil {
		return Relation{}, fmt.Errorf("Erreur lors de la requête : %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Relation{}, fmt.Errorf("Erreur lors de la lecture de la réponse : %v", err)
	}

	var fetchedRelation Relation
	err = json.Unmarshal(body, &fetchedRelation)
	if err != nil {
		return Relation{}, fmt.Errorf("Erreur lors du décodage JSON : %v\nRéponse reçue : %s", err, string(body))
	}

	return fetchedRelation, nil
}
