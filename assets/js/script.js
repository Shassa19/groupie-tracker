const handle = document.querySelector('.scroll-handle');
const bar = document.querySelector('.scroll-bar');

let isDragging = false;
let startX, startLeft;

handle.addEventListener('mousedown', (e) => {
    isDragging = true;
    startX = e.clientX;
    startLeft = handle.offsetLeft;
    handle.style.cursor = 'grabbing';
});

document.addEventListener('mousemove', (e) => {
    if (!isDragging) return;
    let moveX = e.clientX - startX;
    let newLeft = startLeft + moveX;

    // Empêcher de sortir des limites
    if (newLeft < 2) newLeft = 2;
    if (newLeft > bar.offsetWidth - handle.offsetWidth - 2) { 
        newLeft = bar.offsetWidth - handle.offsetWidth - 2; /* Ajoute 1px de marge */
    }
    

    handle.style.left = newLeft + 'px';
});

document.addEventListener('mouseup', () => {
    isDragging = false;
    handle.style.cursor = 'grab';
});


/** Mode nuit */

const toggleSwitch = document.querySelector('.toggle-mode');

toggleSwitch.addEventListener('click', (event) => {
    event.stopPropagation(); // Empêche les conflits de clics
    toggleSwitch.classList.toggle('active');
    document.body.classList.toggle('light-mode');
});


