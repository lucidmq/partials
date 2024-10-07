package heroForm

import "log"

const DATA_KEY = "heroForm"
const CONVERTER_FUNCTION_NAME = "cast_to_heroForm"

type HeroForm struct {
	Title             string
	Subtitle          string
	HoverText         string
	FormLink          string
	ButtonTitle       string
	AuthorityTitle    string
	AuthoritySubtitle string
	AuthorityNumber   string
}

func Cast_to_heroForm(mapper map[string]interface{}) HeroForm {
	var v HeroForm
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(HeroForm)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
