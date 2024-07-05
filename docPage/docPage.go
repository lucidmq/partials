package docPage

import (
	"html/template"
	"log"
)

const DATA_KEY = "docPage"
const CONVERTER_FUNCTION_NAME = "cast_to_doc_page"

type SectionLink struct {
	Title string
	Link  string
}

type TableOfContents struct {
	SectionTitle string
	SubSections  []string
}

type DocPage struct {
	TOC         []TableOfContents
	OnThisPage  []TableOfContents
	DocHtmlData template.HTML
	Section     string
	Title       string
}

func Cast_to_doc_page(mapper map[string]interface{}) DocPage {
	var v DocPage
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(DocPage)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
