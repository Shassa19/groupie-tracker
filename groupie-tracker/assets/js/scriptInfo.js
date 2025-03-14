// Fonction pour récupérer un cookie par son nom
function getCookie(name) {
    const cookies = document.cookie.split("; ");
    for (let cookie of cookies) {
        let [key, value] = cookie.split("=");
        if (key == name) {
            return decodeURIComponent(value);
        }
    }
    return null;
}

// Fonction pour enregistrer les favoris dans les cookies
function setFavorites(favorites) {
    document.cookie = `favorites=${encodeURIComponent(JSON.stringify(favorites))}; path=/; max-age=31536000`; // On créer notre coockie, on va 
    // donner la liste favorites et la converttir en JSOn pour l'utiliserr, puis on l'encode pour eviter les problemes de caractères spéciaux
    console.log("Cookie 'favorites' set:", document.cookie); // Affiche les cookies dans la console
}

// Fonction pour récupérer les favoris depuis les cookies
function getFavorites() {
    const favs = getCookie("favorites"); // on va recuperer la liste des favoris
    console.log("Cookie 'favorites' retrieved:", favs); // Affiche le contenu du cookie dans la console
    return favs ? JSON.parse(favs) : []; // Si coockie existe on le converti en tableau
}


// Programme principale 

let listeFavoris = getFavorites();

//Recupere l'id de l'artist
const lastPart = window.location.pathname.split("/").pop();
console.log("ID de l'artiste:", lastPart); // Vérifiez si l'ID est bien récupéré

if (isNaN(lastPart)) {
    console.error("ID d'artiste invalide:", lastPart);
}

document.addEventListener("DOMContentLoaded", () => {
    const favoriteBtn = document.querySelector(".favorite-btn");

    // Vérifier si la page actuelle est dans les favoris et mettre à jour l'affichage
    if (listeFavoris.includes(lastPart)) {
        favoriteBtn.classList.add("active"); // Ajoute la classe pour mettre le cœur en rouge
    }

    favoriteBtn.addEventListener("click", function () {
        if (!listeFavoris.includes(lastPart)) {
            listeFavoris.push(lastPart);
            this.classList.add("active");
        } else {
            listeFavoris = listeFavoris.filter(item => item !== lastPart);
            this.classList.remove("active");
        }

        // Actualise la liste des favoris
        setFavorites(listeFavoris);

    });
});
