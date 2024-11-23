package cta1

import "log"

const DATA_KEY = "cta1"
const CONVERTER_FUNCTION_NAME = "cast_to_cta1"

type Cta1 struct {
	Title       string
	Description string
	Button1Text string
	Button1Link string
	Button2Text string
	Button2Link string
}

func Cast_to_cta1(mapper map[string]interface{}) Cta1 {
	var v Cta1
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Cta1)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() Cta1 {
	return Cta1{
		Title:       "Streamline Your Workflow, Amplify Your Success",
		Description: "Discover the all-in-one platform designed to simplify tasks, boost productivity, and scale your business effortlessly. From collaboration to analytics, we've got you covered—anytime, anywhere.",
		Button1Text: "Get Started",
		Button1Link: "#",
		Button2Text: "Learn More",
		Button2Link: "#",
	}
}
