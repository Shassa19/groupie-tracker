package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1"

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
