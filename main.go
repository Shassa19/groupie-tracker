package main

import (
	"fmt"
	api "groupie-tracker/Api"
	apiRelation "groupie-tracker/ApiRelation"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var artists []api.Artist
var relation apiRelation.Relation

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists) // On donne les artistes à la page HTML
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID depuis l'URL après "/infoartist/
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
	relation, _ = apiRelation.FetchInfos(id)

	// Charge et affiche la page infoartist.html
	tmpl, err := template.ParseFiles("infoArtist.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}

	println(relation.DatesLocations["georgia-usa"])

	type Data struct {
		Artist   api.Artist
		Relation apiRelation.Relation
	}
	data := Data{Artist: artists[id-1], Relation: relation}
	println(data.Relation.DatesLocations)

	tmpl.Execute(w, data) // Passe l'artiste à la page HTML
}

func main() {
	// URL de l'API principale pour récupérer les artistes
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// On récupère les artistes depuis l'API principale via la fonction fetchArtists du package api
	var err error
	artists, err = api.FetchArtists(url)
	if err != nil {
		fmt.Println(err) // Affichage d'une erreur si la récupération échoue
		return
	}

	// On démarre le serveur HTTP
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/infoartist/", artistHandler)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
