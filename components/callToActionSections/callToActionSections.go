package calltoactionsections

import (
	"partials/components/callToActionSections/cta1"
	"partials/components/callToActionSections/cta2"
	"partials/components/callToActionSections/emailInputCTA"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		cta1.CONVERTER_FUNCTION_NAME:          cta1.Cast_to_cta1,
		cta2.CONVERTER_FUNCTION_NAME:          cta2.Cast_to_cta2,
		emailInputCTA.CONVERTER_FUNCTION_NAME: emailInputCTA.Cast_to_emailInputCTA,
	}
	return funcMap
}
