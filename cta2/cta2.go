package cta2

import "log"

const DATA_KEY = "cta2"
const CONVERTER_FUNCTION_NAME = "cast_to_cta2"

type Cta2 struct {
	Title       string
	Description string
	ButtonText  string
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
