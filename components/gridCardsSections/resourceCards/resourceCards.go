package resourceCards

import (
	"log"
)

const DATA_KEY = "resourceCards"
const CONVERTER_FUNCTION_NAME = "cast_to_resource_cards"

type ResourceCard struct {
	Title                string
	Link                 string
	Section1Subtitle     string
	Section1Description  string
	Section2Subtitle     string
	Section2FeaturesList []string
}

func Cast_to_resource_cards(mapper map[string]interface{}) []ResourceCard {
	var v []ResourceCard
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.([]ResourceCard)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() []ResourceCard {
	featuresList := []string{
		"Keyword Research Template",
		"On-Page SEO Template",
		"SEO Strategy Template",
		"SEO Marketing Plan Template",
		"SEO Proposal Template",
		"Monthly SEO Report Template",
		"SEO Audit Template",
	}

	resourceCards := []ResourceCard{
		{
			Title:                "My guide",
			Link:                 "#",
			Section1Subtitle:     "What You'll Learn",
			Section1Description:  "Get access to priceless SEO templates in this section. With tools to help with every stage of an SEO campaign. Including keyword research, on-page SEO, monthly reporting... and even a client proposal with service contract.",
			Section2Subtitle:     "7 Resources",
			Section2FeaturesList: featuresList,
		},
		{
			Title:                "My guide",
			Link:                 "#",
			Section1Subtitle:     "What You'll Learn",
			Section1Description:  "Get access to priceless SEO templates in this section. With tools to help with every stage of an SEO campaign. Including keyword research, on-page SEO, monthly reporting... and even a client proposal with service contract.",
			Section2Subtitle:     "7 Resources",
			Section2FeaturesList: featuresList,
		},
	}
	return resourceCards
}
