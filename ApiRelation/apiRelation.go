package apiRelation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// On définit la structure de l'Artist
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var relation Relation

// Fonction pour récupérer les données depuis l'API
func FetchInfos(id int) (Relation, error) {
	println(id)
	url := "https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id)
	println(url)
	response, err1 := http.Get(url)

	// On récupère les relations depuis l'API relation via la fonction fetcInfos du package apiRelation
	if err1 != nil {
		fmt.Println("Erreur lors de la récupération des relations :", err1)
	} else {
		fmt.Println("Relation récupérée :")
	}

	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		return Relation{}, fmt.Errorf("Erreur lors de la lecture de la réponse : %v", err2)
	}

	var fetchedRelation Relation
	err3 := json.Unmarshal(body, &fetchedRelation)
	println("aaa", fetchedRelation.ID)
	if err3 != nil {
		return Relation{}, fmt.Errorf("Erreur lors du décodage JSON : %v\nRéponse reçue : %s", err3, string(body))
	}

	return fetchedRelation, nil
}
