function selectFeature(node, evt, featureName) {
  var i, tablinks, otherTabLinks, imageSections;
  tablinks = document.getElementsByClassName("tab-button-selected");
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace("tab-button-selected", "tab-button");
  }
  otherTabLinks = document.getElementsByClassName("tab-container-selected");
  for (i = 0; i < otherTabLinks.length; i++) {
    otherTabLinks[i].className = otherTabLinks[i].className.replace("tab-container-selected", "tab-container");
  }
  imageSections = document.getElementsByClassName("placeholderClass")
  for (i = 0; i < imageSections.length; i++) {
    imageSections[i].style.display = "none";
  }

  evt.currentTarget.className = "tab-button-selected";
  node.parentNode.parentNode.className = "tab-container-selected";
  document.getElementById(featureName).style.display = "block"
}