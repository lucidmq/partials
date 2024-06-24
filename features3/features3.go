package features3

import "log"

const DATA_KEY = "features3"
const CONVERTER_FUNCTION_NAME = "cast_to_features3"

type Step struct {
	Title       string
	Description string
	SvgString   string
}

type IntegrationHow struct {
	Title       string
	Subtitle    string
	Description string
	Steps       []Step
}

func Cast_to_features3(mapper map[string]interface{}) IntegrationHow {
	var v IntegrationHow
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(IntegrationHow)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
