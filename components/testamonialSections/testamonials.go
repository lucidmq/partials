package testamonials

import (
	"partials/components/testamonialSections/testamonialCards"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		testamonialCards.CONVERTER_FUNCTION_NAME: testamonialCards.Cast_to_testamonialCards,
	}
	return funcMap
}
