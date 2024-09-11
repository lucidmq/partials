package integrationHero

import "log"

const DATA_KEY = "integrationHeroData"
const CONVERTER_FUNCTION_NAME = "cast_to_integration_hero"

type IntegrationHero struct {
	App1              string
	App2              string
	App1ImageLocation string
	App2ImageLocation string
}

func Cast_to_integration_hero(mapper map[string]interface{}) IntegrationHero {
	var v IntegrationHero
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(IntegrationHero)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
