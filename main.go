package main

import (
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

	tmpl, err := template.ParseFiles("infoartist.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artist)
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// on utilise fetchArtist pour récuperer les artists depuis l'API
	var err error
	artists, err = api.FetchArtists(url)
	if err != nil {
		log.Printf("Erreur lors de la récupération des artistes depuis l'API : %v", err)
		return
	}

	// Affichage des artistes dans la console avec la fonction displayArtists du package api
	api.DisplayArtists(artists)

	// On démarre le serveur http://localhost:8080
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/infoartist/", artistHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
