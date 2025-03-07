package gridCardsSections

import (
	"html/template"
	"partials/components/gridCardsSections/resourceCards"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		resourceCards.CONVERTER_FUNCTION_NAME: resourceCards.Cast_to_resource_cards,
	}
	return funcMap
}
