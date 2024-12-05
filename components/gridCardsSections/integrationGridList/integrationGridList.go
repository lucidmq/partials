package integrationGridList

import "log"

const DATA_KEY = "integrationGridListData"
const CONVERTER_FUNCTION_NAME = "cast_to_integration_grid_list"

type Integration struct {
	App1              string
	App2              string
	App1ImageLocation string
	App2ImageLocation string
	Link              string
}

func Cast_to_integration_grid_list(mapper map[string]interface{}) []Integration {
	var v []Integration
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.([]Integration)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
