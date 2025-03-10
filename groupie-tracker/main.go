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
var artistsEntier []api.Artist
var relation apiRelation.Relation
var filtre string

type Favorites []string

func filter(listeArtist []api.Artist, filtre string) []api.Artist {
	var listeFiltre []api.Artist

	if filtre == "creation_avant_2000" {
		for _, artist := range listeArtist {
			if artist.CreationDate < 2000 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
	}

	if filtre == "creation_apres_2000" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 2000 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
	}

	if filtre == "solo" {
		for _, artist := range listeArtist {
			if len(artist.Members) == 1 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
	}

	if filtre == "effacer_filtre" {
		listeFiltre = artistsEntier
	}

	return listeFiltre
}

func filtrerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}

	// Vérifier si c'est une requête POST
	if r.Method == http.MethodPost {

		filtre := r.FormValue("filter")

		// Appliquer le filtre à la liste des artistes
		artists = filter(artists, filtre)
	}

	tmpl.Execute(w, artists)
}

// Handler pour afficher la page HTML
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Charger le template HTML
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)

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

	artistsEntier = artists

	// On démarre le serveur HTTP
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/filtre", filtrerHandler)
	http.HandleFunc("/infoartist/", artistHandler)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
