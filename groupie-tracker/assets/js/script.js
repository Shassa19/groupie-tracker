/** Mode nuit */

const toggleSwitch = document.querySelector('.toggle-mode');

toggleSwitch.addEventListener('click', () => {
    toggleSwitch.classList.toggle('active');
    document.body.classList.toggle('light-mode');
});