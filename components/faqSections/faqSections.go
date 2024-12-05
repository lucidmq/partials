package faqsections

import (
	"html/template"
	"partials/components/faqSections/newFaq"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		newFaq.DATA_KEY: newFaq.Cast_to_faq,
	}
	return funcMap
}
