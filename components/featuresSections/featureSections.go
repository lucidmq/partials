package featuressections

import (
	"partials/components/featuresSections/feature1"
	"partials/components/featuresSections/feature2"
	"partials/components/featuresSections/featuresSimple"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		feature1.CONVERTER_FUNCTION_NAME:       feature1.Cast_to_feature1,
		feature2.CONVERTER_FUNCTION_NAME:       feature2.Cast_to_feature2,
		featuresSimple.CONVERTER_FUNCTION_NAME: featuresSimple.Cast_to_featuresSimple,
	}
	return funcMap
}
