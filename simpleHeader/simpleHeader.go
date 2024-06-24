package simpleHeader

import "log"

const DATA_KEY = "simpleHeaderData"
const CONVERTER_FUNCTION_NAME = "cast_to_simple_header"

type SimpleHeader struct {
	Title       string
	Description string
}

func Cast_to_simple_header(mapper map[string]interface{}) SimpleHeader {
	var v SimpleHeader
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(SimpleHeader)
		if !ok {
			log.Fatal("Unable to cast....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
