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