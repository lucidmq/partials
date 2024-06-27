package profileHeader

import "log"

const DATA_KEY = "profileHeader"
const CONVERTER_FUNCTION_NAME = "cast_to_profile_header"

type ProfileHeader struct {
	Title           string
	Description     string
	BackgroundImage string
	ProfileImage    string
}

func Cast_to_profile_header(mapper map[string]interface{}) ProfileHeader {
	var v ProfileHeader
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(ProfileHeader)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
