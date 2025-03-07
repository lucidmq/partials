package simpleHero

import "log"

const DATA_KEY = "simpleHero"
const CONVERTER_FUNCTION_NAME = "cast_to_simple_hero"

type SimpleHero struct {
	Title        string
	Subtitle     string
	Button1Copy  string
	Button1Link  string
	Button2Copy  string
	Button2Link  string
	AlertTopic   string
	AlertContent string
	AlertLink    string
}

func Cast_to_simple_hero(mapper map[string]interface{}) SimpleHero {
	var v SimpleHero
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(SimpleHero)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() SimpleHero {
	return SimpleHero{
		Title:        "Streamline Your Workflow",
		Subtitle:     "Discover the platform designed to simplify tasks and boost productivity",
		Button1Copy:  "Get Started",
		Button1Link:  "#",
		Button2Copy:  "Learn More",
		Button2Link:  "#",
		AlertTopic:   "New!",
		AlertContent: "Collaboration Analytics Suite",
		AlertLink:    "#",
	}
}
