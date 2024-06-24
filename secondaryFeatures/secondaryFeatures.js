function selectSecondaryFeature(node, evt, featureName) {
    const activeTag = 'active-'
    // Get the active element and change it to inactive element
    let elements = document.getElementsByClassName('active-desktop-feature-container')
    let element = elements[0]
    for (const child of element.children) {
        if (child.className.includes(activeTag)) {
            let newClassName = child.className.slice(activeTag.length)
            child.className = newClassName
        }
    }
    element.className = 'desktop-feature-container'

    let selected_element = node.parentNode.parentNode;
    for (const child of selected_element.children) {
        if (!child.className.includes(activeTag)) {
            let newClassName = activeTag + child.className
            child.className = newClassName;
        }
    }
    selected_element.className = 'active-desktop-feature-container';


    let image_element_placement_css = ""
    if (featureName.includes("1")) {
        image_element_placement_css = "translateX(-0%)";
    } else if (featureName.includes("2")) {
        image_element_placement_css = "translateX(-75%)";
    }
    // } else if (featureName.includes("3")) {
    //     image_element_placement_css = "translateX(-200%)";
    // }

    let active_image_elements = document.getElementsByClassName('active-desktop-img-container');
    active_image_elements[0].className = 'desktop-img-container';
    let image_elements = document.getElementsByClassName('desktop-img-container');
    for (image_element of image_elements) {
        image_element.style.transform = image_element_placement_css;
    }

    document.getElementById(featureName).className = 'active-desktop-img-container';
    
}