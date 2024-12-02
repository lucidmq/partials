package testamonials

import "log"

// TODO: this needs to be hooked into the actual template

const DATA_KEY = "testamonials"
const CONVERTER_FUNCTION_NAME = "cast_to_testamonials"

type Testamonials struct {
	Title            string
	Description      string
	TestamonialsInfo []Testamonial
}

type Testamonial struct {
	Comment string
	Name    string
	Role    string
	Image   string
}

func Cast_to_testamonials(mapper map[string]interface{}) Testamonials {
	var v Testamonials
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Testamonials)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
