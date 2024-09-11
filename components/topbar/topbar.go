package topbar

import (
	"html/template"
	"log"
)

const DATA_KEY = "topbar"
const CONVERTER_FUNCTION_NAME = "cast_to_topbar"

type MainLink struct {
	Title string
	Link  string
}

type Topbar struct {
	MainLinks      []MainLink
	LogoSVG        template.HTML //We insert the svg directly into the page
	ShowLoginLinks bool
}

func Cast_to_topbar(mapper map[string]interface{}) Topbar {
	var v Topbar
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Topbar)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
