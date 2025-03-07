package heroComponents

import (
	"partials/components/heroComponents/heroForm"
	"partials/components/heroComponents/heroImageButton"
	"partials/components/heroComponents/integrationHero"
	"partials/components/heroComponents/simpleHero"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		heroForm.CONVERTER_FUNCTION_NAME:        heroForm.Cast_to_heroForm,
		integrationHero.CONVERTER_FUNCTION_NAME: integrationHero.Cast_to_integration_hero,
		simpleHero.CONVERTER_FUNCTION_NAME:      simpleHero.Cast_to_simple_hero,
		heroImageButton.CONVERTER_FUNCTION_NAME: heroImageButton.Cast_to_heroImageButton,
	}
	return funcMap
}
