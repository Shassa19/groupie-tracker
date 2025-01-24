package main

import (
	"fmt"
	api "groupie/Api"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	// Récupère l'ID depuis l'URL après "/infoartist/"
	idParam := strings.TrimPrefix(r.URL.Path, "/infoartist/")
	if idParam == "" {
		http.Error(w, "ID de l'artiste manquant", http.StatusBadRequest)
		return
	}

	// Conversion de l'ID en entier
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 || id > len(artists) {
		http.Error(w, "Artiste non trouvé", http.StatusNotFound)
		return
	}

	// Récupère l'artiste correspondant
	artist := artists[id-1]

	// Charge et affiche la page infoartist.html
	tmpl, err := template.ParseFiles("infoartist.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artist) // Passe l'artiste à la page HTML
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
	http.HandleFunc("/infoartist/", artistHandler)
	http.Handle("/Styles/style.css", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
