package emailInputCTA

import "log"

const DATA_KEY = "emailInputCTA"
const CONVERTER_FUNCTION_NAME = "cast_to_emailInputCTA"

type EmailInputCTA struct {
	Title      string
	Subtitle   string
	Link       string
	ButtonText string
}

func Cast_to_emailInputCTA(mapper map[string]interface{}) EmailInputCTA {
	var v EmailInputCTA
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(EmailInputCTA)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
