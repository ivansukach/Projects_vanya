let slides = document.querySelectorAll('#slides .slide');
let currentSlide = 0;
let slideInterval = setInterval(nextSlide, 4000);

function nextSlide() {
    goToSlide(currentSlide + 1);
}

function previousSlide() {
    goToSlide(currentSlide - 1);
}

function goToSlide(n) {
    slides[currentSlide].className = 'slide';
    currentSlide = (n + slides.length) % slides.length;
    slides[currentSlide].className = 'slide showing';
}

let playing = true;

function pauseSlideshow() {
    playing = false;
    clearInterval(slideInterval);
}

let next = document.getElementById('next');
let previous = document.getElementById('previous');
next.onclick = function() {
    pauseSlideshow();
    nextSlide();
};
previous.onclick = function() {
    pauseSlideshow();
    previousSlide();
};