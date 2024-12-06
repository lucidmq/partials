package aboutSection

import (
	"log"
)

const DATA_KEY = "aboutSection"
const CONVERTER_FUNCTION_NAME = "cast_to_about_section"

type AboutSection struct {
	Quote           string
	Subtitle        string
	Title           string
	ContentSections []string
}

func Cast_to_about_section(mapper map[string]interface{}) AboutSection {
	var v AboutSection
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(AboutSection)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() AboutSection {
	return AboutSection{
		Quote:           "Discover the platform designed to simplify tasks, boost productivity, and scale your business effortlessly.",
		Subtitle:        "Streamline Your Workflow",
		Title:           "Streamline Your Workflow",
		ContentSections: []string{"Discover the platform designed to simplify tasks, boost productivity, and scale your business effortlessly.", "Discover the platform designed to simplify tasks, boost productivity, and scale your business effortlessly.", "Discover the platform designed to simplify tasks, boost productivity, and scale your business effortlessly."},
	}
}
