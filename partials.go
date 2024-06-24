package partials

import (
	"partials/blogArticle"
	"partials/cta2"
	"partials/feature1"
	"partials/features3"
	"partials/integrationCards"
	"partials/integrationGridList"
	"partials/integrationHero"
	"partials/newFaq"
	"partials/pricingCards"
	"partials/simpleHeader"
	"partials/simpleHero"
	"partials/singleColumnBlog"
	"text/template"
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		newFaq.CONVERTER_FUNCTION_NAME:              newFaq.Cast_to_faq,
		simpleHeader.CONVERTER_FUNCTION_NAME:        simpleHeader.Cast_to_simple_header,
		integrationGridList.CONVERTER_FUNCTION_NAME: integrationGridList.Cast_to_integration_grid_list,
		integrationCards.CONVERTER_FUNCTION_NAME:    integrationCards.Cast_to_integration_Cards,
		features3.CONVERTER_FUNCTION_NAME:           features3.Cast_to_features3,
		integrationHero.CONVERTER_FUNCTION_NAME:     integrationHero.Cast_to_integration_hero,
		singleColumnBlog.CONVERTER_FUNCTION_NAME:    singleColumnBlog.Cast_to_single_col_blog,
		blogArticle.CONVERTER_FUNCTION_NAME:         blogArticle.Cast_to_blog_article,
		simpleHero.CONVERTER_FUNCTION_NAME:          simpleHero.Cast_to_simple_hero,
		cta2.CONVERTER_FUNCTION_NAME:                cta2.Cast_to_cta2,
		feature1.CONVERTER_FUNCTION_NAME:            feature1.Cast_to_feature1,
		pricingCards.CONVERTER_FUNCTION_NAME:        pricingCards.Cast_to_pricing_cards,
	}

	return funcMap
}
