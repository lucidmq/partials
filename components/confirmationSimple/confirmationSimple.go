package confirmationSimple

import "log"

const DATA_KEY = "confirmationSimple"
const CONVERTER_FUNCTION_NAME = "cast_to_confirmationSimple"

type ConfirmationSimple struct {
	Title       string
	Description string
	Button1Text string
	Button1Link string
}

func Cast_to_confirmationSimple(mapper map[string]interface{}) ConfirmationSimple {
	var v ConfirmationSimple
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(ConfirmationSimple)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
