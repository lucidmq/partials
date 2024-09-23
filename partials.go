package partials

import (
	"partials/components/blogArticle"
	"partials/components/cta1"
	"partials/components/cta2"
	"partials/components/docPage"
	"partials/components/feature1"
	"partials/components/features3"
	"partials/components/footer"
	"partials/components/fromTheBlogSection"
	"partials/components/integrationCards"
	"partials/components/integrationGridList"
	"partials/components/integrationHero"
	"partials/components/linker"
	"partials/components/logoCloudSimple"
	"partials/components/myProductItems"
	"partials/components/newFaq"
	"partials/components/pricingCards"
	"partials/components/product"
	"partials/components/profileHeader"
	"partials/components/simpleHeader"
	"partials/components/simpleHero"
	"partials/components/singleColumnBlog"
	"partials/components/topbar"
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
		cta1.CONVERTER_FUNCTION_NAME:                cta1.Cast_to_cta1,
		docPage.CONVERTER_FUNCTION_NAME:             docPage.Cast_to_doc_page,
		profileHeader.CONVERTER_FUNCTION_NAME:       profileHeader.Cast_to_profile_header,
		topbar.CONVERTER_FUNCTION_NAME:              topbar.Cast_to_topbar,
		fromTheBlogSection.CONVERTER_FUNCTION_NAME:  fromTheBlogSection.Cast_to_from_the_blog_section,
		linker.CONVERTER_FUNCTION_NAME:              linker.Cast_to_linker,
		logoCloudSimple.CONVERTER_FUNCTION_NAME:     logoCloudSimple.Cast_to_tester,
		footer.CONVERTER_FUNCTION_NAME:              footer.Cast_to_footer,
		product.CONVERTER_FUNCTION_NAME:             product.Cast_to_product,
		myProductItems.CONVERTER_FUNCTION_NAME:      myProductItems.Cast_to_myProductItems,
		"add":                                       Add,
		"sub":                                       Sub,
	}

	return funcMap
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
