package main

import (
	"encoding/json"
	"fmt"
	api "groupie-tracker/Api"
	apiRelation "groupie-tracker/ApiRelation"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
)

var artists []api.Artist
var artistsEntier []api.Artist
var relation apiRelation.Relation
var filtre string

type Favorites []string

// Fonction pour récupérer un cookie par son nom
func getCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// Fonction pour récupérer les favoris depuis les cookies et les convertir en liste d'entiers
func getFavorites(r *http.Request) ([]int, error) {
	// Récupérer le cookie "favorites"
	cookieValue, err := getCookie(r, "favorites")
	if err != nil {
		return nil, fmt.Errorf("cookie not found")
	}

	// Décoder l'URL pour obtenir la vraie valeur JSON
	decodedValue, err := url.QueryUnescape(cookieValue)
	if err != nil {
		fmt.Println("Erreur lors du décodage de l'URL:", err)
		return nil, fmt.Errorf("failed to decode cookie data")
	}

	// Décoder la chaîne JSON en un tableau de chaînes
	var favoritesStr []string
	err = json.Unmarshal([]byte(decodedValue), &favoritesStr)
	if err != nil {
		fmt.Println("Erreur lors de la conversion du cookie en JSON:", err)
		return nil, fmt.Errorf("failed to parse cookie data")
	}

	// Convertir les chaînes en entiers
	var favorites []int
	for _, str := range favoritesStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en entier:", err)
			return nil, fmt.Errorf("failed to convert cookie data to int")
		}
		favorites = append(favorites, num)
	}

	return favorites, nil
}

func filter(r *http.Request, listeArtist []api.Artist, filtre string) []api.Artist {
	var listeFiltre []api.Artist

	if filtre == "creation_entre_1950_1959" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 1950 && artist.CreationDate <= 1959 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "creation_entre_1960_1969" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 1960 && artist.CreationDate <= 1969 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "creation_entre_1970_1979" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 1970 && artist.CreationDate <= 1979 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "creation_entre_1980_1989" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 1980 && artist.CreationDate <= 1989 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "creation_entre_1990_1999" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 1990 && artist.CreationDate <= 1999 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "creation_entre_2000_2009" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 2000 && artist.CreationDate <= 2009 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "creation_entre_2010_2019" {
		for _, artist := range listeArtist {
			if artist.CreationDate >= 2010 && artist.CreationDate <= 2019 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "un_membre" {
		for _, artist := range listeArtist {
			if len(artist.Members) == 1 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "deux-membres" {
		for _, artist := range listeArtist {
			if len(artist.Members) == 2 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "groupe" {
		for _, artist := range listeArtist {
			if len(artist.Members) > 2 {
				listeFiltre = append(listeFiltre, artist)
			}
		}
		if len(listeFiltre) == 0 {
			return filter(r, artistsEntier, filtre)
		}
	}

	if filtre == "favoris" {
		// Récupérer les IDs des favoris à partir du cookie
		favorites, err := getFavorites(r)
		if err != nil {
			// Si aucune erreur, continuer avec une liste vide de favoris
			fmt.Println("Erreur lors de la récupération des favoris:", err)
			return nil
		}

		// Filtrer la liste des artistes pour ne garder que ceux qui sont dans les favoris
		for _, artist := range listeArtist {
			for _, favID := range favorites {
				if strconv.Itoa(artist.ID) == strconv.Itoa(favID) {
					listeFiltre = append(listeFiltre, artist)
				}
			}
		}
	}

	if filtre == "effacer_filtre" {
		listeFiltre = artistsEntier
	}

	return listeFiltre
}

func filtrerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		filtre := r.FormValue("filter")

		// Applique le filtre
		artists = filter(r, artists, filtre)
	} else {
		artists = artistsEntier
	}

	// Charge et affiche le template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)
}

// Handler pour afficher la page HTML
func homeHandler(w http.ResponseWriter, r *http.Request) {

	if len(artists) == 0 {
		artists = artistsEntier // Utilise la liste complète par défaut
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)

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
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Recherche l'artiste correspondant dans la liste complète `artistsEntier`
	var selectedArtist api.Artist
	found := false

	for _, artist := range artistsEntier { // On utilise la liste complète d'artistes, pas filtrée
		if artist.ID == id {
			selectedArtist = artist
			found = true
			break
		}
	}

	// Si l'artiste n'est pas trouvé
	if !found {
		http.Error(w, "Artiste non trouvé", http.StatusNotFound)
		return
	}

	// Récupère les informations liées à l'artiste
	relation, _ = apiRelation.FetchInfos(id)

	// Charge et affiche la page infoartist.html
	tmpl, err := template.ParseFiles("infoArtist.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page HTML", http.StatusInternalServerError)
		return
	}

	// Structure des données à passer à la page HTML
	type Data struct {
		Artist   api.Artist
		Relation apiRelation.Relation
	}

	// Création des données à passer au template
	data := Data{Artist: selectedArtist, Relation: relation}

	// Affichage du template avec les données
	tmpl.Execute(w, data)
}

// Artistes format JSON
func apiArtistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(artists)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage JSON", http.StatusInternalServerError)
	}
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
	http.HandleFunc("/api/artists", apiArtistsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
