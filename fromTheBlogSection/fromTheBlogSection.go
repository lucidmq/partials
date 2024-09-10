package fromTheBlogSection

import (
	"log"
	"partials/singleColumnBlog"
)

const DATA_KEY = "fromTheBlogSection"
const CONVERTER_FUNCTION_NAME = "cast_to_from_the_blog_section"

type FromTheBlogSection struct {
	PrimaryPost singleColumnBlog.BlogPageDetails
	OtherPosts  []singleColumnBlog.BlogPageDetails
}

func Cast_to_from_the_blog_section(mapper map[string]interface{}) FromTheBlogSection {
	var v FromTheBlogSection
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(FromTheBlogSection)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
