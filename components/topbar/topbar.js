function toggleMenu(event) {
    var x = document.getElementById("resNav");
    if (x.className === "mobile_sub_object") {
      x.className = "mobile_sub_object_responsive";
    } else {
      x.className = "mobile_sub_object";
    }
}
