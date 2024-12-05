package integrationCards

import "log"

const DATA_KEY = "integrationCards"
const CONVERTER_FUNCTION_NAME = "cast_to_integration_cards"

type ExampleIntegration struct {
	Description       string
	App1ImageLocation string
	App2ImageLocation string
}

type IntegrationFeatures struct {
	Title              string
	Description        string
	ExampleIntegations []ExampleIntegration
}

func Cast_to_integration_Cards(mapper map[string]interface{}) IntegrationFeatures {
	var v IntegrationFeatures
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(IntegrationFeatures)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
