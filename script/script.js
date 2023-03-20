function openPopUp(element) {
    element.nextElementSibling.style.display = "block";
    document.body.style.overflow = "hidden";
}

function closePopUp(element) {
    element.parentElement.parentElement.style.display = "none";
    document.body.style.overflow = "auto";
}

function showSlides1(element) {
    element = element.parentElement
    element.nextElementSibling.style.display = "block"
    element.nextElementSibling.nextElementSibling.style.display = "none"
    element.nextElementSibling.nextElementSibling.nextElementSibling.style.display = "none"
}

function showSlides2(element) {
    element = element.parentElement
    element.nextElementSibling.style.display = "none"
    element.nextElementSibling.nextElementSibling.style.display = "block"
    element.nextElementSibling.nextElementSibling.nextElementSibling.style.display = "none"
}

function showSlides3(element) {
    element = element.parentElement
    element.nextElementSibling.style.display = "none"
    element.nextElementSibling.nextElementSibling.style.display = "none"
    element.nextElementSibling.nextElementSibling.nextElementSibling.style.display = "block"
}