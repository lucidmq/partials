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

func NewDummy() Feature1 {
	return Feature1{
		Title:         "Effortless and Powerful Data Analysis",
		Subtitle:      "Create professional-grade reports in seconds. All reports are fully editable, allowing you to adjust data as your needs evolve.",
		Icon1Title:    "Make Confident Decisions with Clear Insights",
		Icon1Subtitle: "Generate detailed reports quickly—visualize key metrics, refine your analysis, and stay in control of your outcomes.",
		Icon2Title:    "User-Friendly Input for Fast Results",
		Icon2Subtitle: "Designed for simplicity—no advanced knowledge required. Just input your details and get accurate results instantly.",
		LinkText:      "Try It Free Today",
		LinkPath:      "#",
	}
}
