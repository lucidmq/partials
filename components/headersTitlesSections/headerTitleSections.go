package headerTitleSections

import (
	"partials/components/headersTitlesSections/simpleHeader"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		simpleHeader.CONVERTER_FUNCTION_NAME: simpleHeader.Cast_to_simple_header,
	}
	return funcMap
}
