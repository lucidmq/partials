package newFaq

import "log"

const DATA_KEY = "faqData"
const CONVERTER_FUNCTION_NAME = "cast_to_faq"

type NewFaq struct {
	Question string
	Answer   string
}

func Cast_to_faq(mapper map[string]interface{}) []NewFaq {
	var v []NewFaq
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.([]NewFaq)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
