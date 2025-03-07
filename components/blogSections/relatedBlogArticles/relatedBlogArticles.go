package relatedBlogArticles

import (
	"log"
	"partials/types"
)

const DATA_KEY = "relatedBlogArticles"
const CONVERTER_FUNCTION_NAME = "cast_to_related_blog_articles"

type RelatedBlogArticles struct {
	Title string
	// Change this to enforce 2
	Posts [2]types.BlogPageDetails
}

func Cast_to_related_blog_articles(mapper map[string]interface{}) RelatedBlogArticles {
	var v RelatedBlogArticles
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(RelatedBlogArticles)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
