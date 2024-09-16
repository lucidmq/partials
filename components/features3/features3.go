package features3

import "log"

const DATA_KEY = "features3"
const CONVERTER_FUNCTION_NAME = "cast_to_features3"

type SubFeature struct {
	Title       string
	Description string
	SvgString   string
}

type Features3 struct {
	Title       string
	Subtitle    string
	Description string
	Steps       []SubFeature
}

func Cast_to_features3(mapper map[string]interface{}) Features3 {
	var v Features3
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Features3)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
