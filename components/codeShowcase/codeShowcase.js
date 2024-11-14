
function changeCodeTab(evt, cityName) {
    evt.preventDefault()
    var i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("SEXJEWKEHZWYJWC");
    for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].className = "SEXJEWKEHZWYJWC_hidden"
    }
    document.getElementById(cityName).className = "SEXJEWKEHZWYJWC";

    tablinks = document.getElementsByClassName("active_code_tab");
    for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = "inactive_code_tab"
    }

    evt.currentTarget.className += " active_code_tab";
}