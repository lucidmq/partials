package linker

import (
	"html/template"
	"log"
)

const DATA_KEY = "linker"
const CONVERTER_FUNCTION_NAME = "cast_to_linker"

type Linker struct {
	Link        string
	ButtonStyle string
	IconSvg     template.HTML //We insert the svg directly into the page
	LinkText    string
}

func Cast_to_linker(mapper map[string]interface{}) []Linker {
	var v []Linker
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.([]Linker)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
