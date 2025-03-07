package heroImageButton

import "log"

const DATA_KEY = "heroImageButton"
const CONVERTER_FUNCTION_NAME = "Cast_to_heroImageButton"

type heroImageButton struct {
	Title       string
	Subtitle    string
	ButtonTitle string
	ButtonLink  string
	ImageLink   string
}

func Cast_to_heroImageButton(mapper map[string]interface{}) heroImageButton {
	var v heroImageButton
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(heroImageButton)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() heroImageButton {
	return heroImageButton{
		Title:       "Payments tool for software companies",
		Subtitle:    "From checkout to global sales tax compliance, companies around the world use Flowbite to simplify their payment stack.",
		ButtonLink:  "#",
		ButtonTitle: "Try Now",
		ImageLink:   "https://flowbite.s3.amazonaws.com/blocks/marketing-ui/hero/phone-mockup.png",
	}
}
