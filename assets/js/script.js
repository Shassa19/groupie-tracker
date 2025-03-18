/** Mode nuit */

const toggleSwitch = document.querySelector('.toggle-mode');

toggleSwitch.addEventListener('click', () => {
    toggleSwitch.classList.toggle('active');
    document.body.classList.toggle('light-mode');
});


const filtreBtn = document.querySelector(".boutonFiltre");
const menuFiltre = document.querySelector(".menu_filtre");

filtreBtn.addEventListener("click", () => {
    menuFiltre.classList.toggle('show');
});