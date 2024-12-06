package navigationComponents

import (
	"partials/components/navigationComponents/footer"
	"partials/components/navigationComponents/topbar"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		footer.CONVERTER_FUNCTION_NAME: footer.Cast_to_footer,
		topbar.CONVERTER_FUNCTION_NAME: topbar.Cast_to_topbar,
	}
	return funcMap
}
