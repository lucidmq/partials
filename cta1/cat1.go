package cta1

import "log"

const DATA_KEY = "cta1"
const CONVERTER_FUNCTION_NAME = "cast_to_cta1"

type Cta1 struct {
	Title       string
	Description string
	Button1Text string
	Button2Text string
}

func Cast_to_cta1(mapper map[string]interface{}) Cta1 {
	var v Cta1
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Cta1)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
