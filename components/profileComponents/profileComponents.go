package profileComponents

import (
	"partials/components/profileComponents/linker"
	"partials/components/profileComponents/myProductItems"
	"partials/components/profileComponents/profileHeader"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		linker.CONVERTER_FUNCTION_NAME:         linker.Cast_to_linker,
		myProductItems.CONVERTER_FUNCTION_NAME: myProductItems.Cast_to_myProductItems,
		profileHeader.CONVERTER_FUNCTION_NAME:  profileHeader.Cast_to_profile_header,
	}
	return funcMap
}
