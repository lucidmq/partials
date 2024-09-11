package feature1

import "log"

const DATA_KEY = "feature1"
const CONVERTER_FUNCTION_NAME = "cast_to_feature1"

type Feature1 struct {
	Title         string
	Subtitle      string
	Icon1Title    string
	Icon1Subtitle string
	Icon2Title    string
	Icon2Subtitle string
	LinkText      string
	LinkPath      string
}

func Cast_to_feature1(mapper map[string]interface{}) Feature1 {
	var v Feature1
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Feature1)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
