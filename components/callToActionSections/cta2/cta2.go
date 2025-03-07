package cta2

import (
	"html/template"
	"log"
)

const DATA_KEY = "cta2"
const CONVERTER_FUNCTION_NAME = "cast_to_cta2"

type Cta2 struct {
	Title               string
	Description         string
	ButtonText          string
	ButtonLink          string
	ButtonFunction      template.JS
	BackgroundImagePath string
}

func Cast_to_cta2(mapper map[string]interface{}) Cta2 {
	var v Cta2
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Cta2)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() Cta2 {
	return Cta2{
		Title:       "Streamline Your Workflow",
		Description: "Discover the platform designed to simplify tasks, boost productivity, and scale your business effortlessly.",
		ButtonText:  "Get Started",
		ButtonLink:  "#",
	}
}
