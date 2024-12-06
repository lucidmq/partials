package feature2

import "log"

const DATA_KEY = "feature2"
const CONVERTER_FUNCTION_NAME = "cast_to_feature2"

type SubFeature struct {
	Title       string
	Description string
	SvgString   string
}

type Feature2 struct {
	Title       string
	Subtitle    string
	Description string
	Steps       []SubFeature
}

func Cast_to_feature2(mapper map[string]interface{}) Feature2 {
	var v Feature2
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Feature2)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

var svgStrings = map[string]string{
	"cloud": "M5.5 17a4.5 4.5 0 01-1.44-8.765 4.5 4.5 0 018.302-3.046 3.5 3.5 0 014.504 4.272A4 4 0 0115 17H5.5zm3.75-2.75a.75.75 0 001.5 0V9.66l1.95 2.1a.75.75 0 101.1-1.02l-3.25-3.5a.75.75 0 00-1.1 0l-3.25 3.5a.75.75 0 101.1 1.02l1.95-2.1v4.59z",
	"lock":  "M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z",
	"cycle": "M15.312 11.424a5.5 5.5 0 01-9.201 2.466l-.312-.311h2.433a.75.75 0 000-1.5H3.989a.75.75 0 00-.75.75v4.242a.75.75 0 001.5 0v-2.43l.31.31a7 7 0 0011.712-3.138.75.75 0 00-1.449-.39zm1.23-3.723a.75.75 0 00.219-.53V2.929a.75.75 0 00-1.5 0V5.36l-.31-.31A7 7 0 003.239 8.188a.75.75 0 101.448.389A5.5 5.5 0 0113.89 6.11l.311.31h-2.432a.75.75 0 000 1.5h4.243a.75.75 0 00.53-.219z",
}

func NewDummy() Feature2 {
	steps := []SubFeature{
		{
			Title:       "Create Mapping",
			Description: "A mapping describes declaratively what information you want pulled from your existing system. Mappings are as flexible as your source application.",
			SvgString:   svgStrings["lock"],
		},
		{
			Title:       "Sync Jobs",
			Description: "A job runs at a defined period. Using the fields specified in the mapping, the job automatically updates existing rows in the spreadsheet with the new data.",
			SvgString:   svgStrings["cycle"],
		},
		{
			Title:       "Customize and Unlock Insights",
			Description: "Start adding dashboards and new reports to your workflows. As data gets synced, your subsequent dashboards and reports will update automatically.",
			SvgString:   svgStrings["cloud"],
		},
	}

	return Feature2{
		Title:       "How it all works",
		Subtitle:    "Integrate faster",
		Description: "Our integration process is effortless and powerful. This is how it workw.",
		Steps:       steps,
	}
}
