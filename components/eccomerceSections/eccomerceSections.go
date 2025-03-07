package eccomerceSections

import (
	"partials/components/eccomerceSections/product"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		product.CONVERTER_FUNCTION_NAME: product.Cast_to_product,
	}
	return funcMap
}
