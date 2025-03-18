document.getElementById('searchInput').addEventListener('input', function() {
    let query = this.value.toLowerCase();
    let resultsContainer = document.getElementById('searchResults');

    if (query.length > 0) {
    // Appel à l'API pour récupérer les artistes
    fetch('/api/artists')
    .then(response => response.json())
    .then(artists => {
    // Filtrage des artistes en fonction du texte tapé
    let filteredArtists = artists.filter(artist =>
    artist.name.toLowerCase().includes(query)
    );

    // Si aucun résultat, on cache le conteneur des résultats
    if (filteredArtists.length === 0) {
    resultsContainer.style.display = 'none';
} else {
    // Affichage des résultats dans le conteneur
    resultsContainer.innerHTML = filteredArtists.map(artist => {
    return `<div onclick="window.location.href='/infoartist/${artist.id}'">
                                    ${artist.name}
                                </div>`;
}).join('');
    resultsContainer.style.display = 'block';
}
})
    .catch(error => {
    console.error('Erreur de récupération des artistes:', error);
    resultsContainer.style.display = 'none';
});
} else {
    resultsContainer.style.display = 'none'; // Si rien n'est tapé, on cache les résultats
}
});