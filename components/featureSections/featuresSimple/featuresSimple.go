package featuresSimple

import "log"

const DATA_KEY = "featuresSimeple"
const CONVERTER_FUNCTION_NAME = "cast_to_featuresSimple"

type FeaturesSimple struct {
	Section1Title     string
	Section1Body      string
	Section1Subtitle  string
	Section1Points    []string
	Section1ImagePath string
	Section2Title     string
	Section2Body      string
	Section2Subtitle  string
	Section2Points    []string
	Section2ImagePath string
}

func Cast_to_featuresSimple(mapper map[string]interface{}) FeaturesSimple {
	var v FeaturesSimple
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(FeaturesSimple)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
