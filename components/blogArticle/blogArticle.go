package blogArticle

import (
	"html/template"
	"log"
	"partials/types"
)

const DATA_KEY = "blogArticle"
const CONVERTER_FUNCTION_NAME = "cast_to_blog_article"

type BlogArticle struct {
	BlogDetails    types.BlogPageDetails
	HtmlDataString template.HTML
}

func Cast_to_blog_article(mapper map[string]interface{}) BlogArticle {
	var v BlogArticle
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(BlogArticle)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
