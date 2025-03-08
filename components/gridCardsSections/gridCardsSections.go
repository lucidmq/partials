package gridCardsSections

import (
	"html/template"
	"partials/components/gridCardsSections/integrationGridList"
	"partials/components/gridCardsSections/resourceCards"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		resourceCards.CONVERTER_FUNCTION_NAME:       resourceCards.Cast_to_resource_cards,
		integrationGridList.CONVERTER_FUNCTION_NAME: integrationGridList.Cast_to_integration_grid_list,
	}
	return funcMap
}
