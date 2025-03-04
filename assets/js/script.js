/** Mode nuit */

const toggleSwitch = document.querySelector('.toggle-mode');

toggleSwitch.addEventListener('click', (event) => {
    event.stopPropagation(); // Empêche les conflits de clics
    toggleSwitch.classList.toggle('active');
    document.body.classList.toggle('light-mode');
});


