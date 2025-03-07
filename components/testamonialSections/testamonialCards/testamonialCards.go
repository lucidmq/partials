package testamonialCards

import "log"

// TODO: this needs to be hooked into the actual template

const DATA_KEY = "testamonialCards"
const CONVERTER_FUNCTION_NAME = "cast_to_testamonialCards"

type TestamonialCards struct {
	Title        string
	Description  string
	Testamonials []TestamonialCard
}

type TestamonialCard struct {
	Comment string
	Name    string
	Role    string
	Image   string
}

func Cast_to_testamonialCards(mapper map[string]interface{}) TestamonialCards {
	var v TestamonialCards
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(TestamonialCards)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() TestamonialCards {
	return TestamonialCards{
		Title:       "thing",
		Description: "thing",
		Testamonials: []TestamonialCard{
			{
				Comment: "heelo",
				Name:    "Stever",
				Role:    "Principal",
				Image:   "/static/images/profile/another.webp",
			},
		},
	}

}
