package blogSections

import (
	"partials/components/blogSections/blogArticle"
	"partials/components/blogSections/fromTheBlogSection"
	"partials/components/blogSections/singleColumnBlog"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		blogArticle.CONVERTER_FUNCTION_NAME:        blogArticle.Cast_to_blog_article,
		singleColumnBlog.CONVERTER_FUNCTION_NAME:   singleColumnBlog.Cast_to_single_col_blog,
		fromTheBlogSection.CONVERTER_FUNCTION_NAME: fromTheBlogSection.Cast_to_from_the_blog_section,
		"add": Add,
		"sub": Sub,
	}
	return funcMap
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
