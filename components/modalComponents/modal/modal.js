// When a link has a #modal in it's href attribute. We want to be able to open the modal.
//This happens on load since we go through all of the link tags
window.onload = function() {
    var anchors = document.getElementsByTagName('a'); // get all <a> tags
    if (anchors) {
        // Use one of the following lines, based on what you require:
        for (let anchor of anchors) {
            if (anchor.getAttribute('href') == "#modal") {
                anchor.addEventListener('click', (e) => openModal(e))
            }
        }
    }
}

function openModal(event) {
    event.preventDefault();
    const bodyNode = document.querySelector("body");
    const template = document.querySelector("#modalTemplate");
    const clone = template.content.cloneNode(true);
    bodyNode.appendChild(clone);
}

function stopClose(event) {
    event.stopPropagation()
}

function removeButton() {
    const element = document.getElementById("modal");
    element.remove()
}