package footer

import (
	"html/template"
	"log"
)

const DATA_KEY = "footer"
const CONVERTER_FUNCTION_NAME = "cast_to_footer"

type MainLink struct {
	Title string
	Link  string
}

type Footer struct {
	SiteName    string
	MainLinks   []MainLink
	LogoSVG     template.HTML //We insert the svg directly into the page
	CompanyName string
}

func Cast_to_footer(mapper map[string]interface{}) Footer {
	var v Footer
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Footer)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Footer data found in the map")
	return v
}
