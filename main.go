package main

import (
	"fmt"
	api "groupie/Api"
	"log"
	"net/http"
	"text/template"
)

var artists []api.Artist

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists) // On donne les artistes à la page HTML
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("infoArtist.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists) // On donne les artistes à la page HTML
}

func main() {
	// URL de l'API pour récupérer les artistes
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// On récupère les artistes depuis l'API via la fonction fetchArtists du package api
	var err error
	artists, err = api.FetchArtists(url)
	if err != nil {
		fmt.Println(err) // Affichage d'une erreur si la récupération échoue
		return
	}

	// Affichage des artistes dans la console avec la fonction displayArtists du package api
	api.DisplayArtists(artists)

	// On démarre le serveur HTTP
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/infoartist", artistHandler)
	http.Handle("/Styles/style.css", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
