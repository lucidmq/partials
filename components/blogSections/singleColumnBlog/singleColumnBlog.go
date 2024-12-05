package singleColumnBlog

import (
	"log"
	"partials/types"
)

const DATA_KEY = "singleColBlog"
const CONVERTER_FUNCTION_NAME = "cast_to_single_col_blog"

type SingleColBlog struct {
	Title       string
	Subtitle    string
	BlogPages   []types.BlogPageDetails
	Pagination  []int
	CurrentPage int
	TotalPages  int
}

func Cast_to_single_col_blog(mapper map[string]interface{}) SingleColBlog {
	var v SingleColBlog
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(SingleColBlog)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
