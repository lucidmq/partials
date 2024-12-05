package logoCloudSimple

import "log"

const DATA_KEY = "logoCloudSimple"
const CONVERTER_FUNCTION_NAME = "cast_to_tester"

type LogoImages struct {
	Alt string
	Src string
}

type LogoCloudSimple struct {
	Title string
	Logos []LogoImages
}

func Cast_to_tester(mapper map[string]interface{}) LogoCloudSimple {
	var v LogoCloudSimple
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(LogoCloudSimple)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
