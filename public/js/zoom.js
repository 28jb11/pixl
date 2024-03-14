document.addEventListener('wheel', event => {
    const grid = document.querySelector('.grid');
    let newSize = parseInt(grid.style.getPropertyValue('--size') || 40);
    let newGap = parseInt(grid.style.getPropertyValue('--gap') || 1);

    if (event.deltaY < 0) { // Zoom in
        newSize += 5;
        newGap += 0.2;
    } else { // Zoom out
        newSize = Math.max(10, newSize - 5); // Prevent the size from becoming too small
        newGap = Math.max(0.2, newGap - 0.2);
    }

    grid.style.setProperty('--size', `${newSize}px`);
    grid.style.setProperty('--gap', `${newGap}px`);
}, {passive: true});
