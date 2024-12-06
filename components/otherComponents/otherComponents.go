package otherComponents

import (
	"partials/components/otherComponents/aboutSection"
	"partials/components/otherComponents/confirmationSimple"
	"partials/components/otherComponents/docPage"
	"partials/components/otherComponents/pricingCards"
	"partials/components/otherComponents/testamonials"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		aboutSection.CONVERTER_FUNCTION_NAME:       aboutSection.Cast_to_about_section,
		confirmationSimple.CONVERTER_FUNCTION_NAME: confirmationSimple.Cast_to_confirmationSimple,
		docPage.CONVERTER_FUNCTION_NAME:            docPage.Cast_to_doc_page,
		pricingCards.CONVERTER_FUNCTION_NAME:       pricingCards.Cast_to_pricing_cards,
		testamonials.CONVERTER_FUNCTION_NAME:       testamonials.Cast_to_testamonials,
	}
	return funcMap
}
