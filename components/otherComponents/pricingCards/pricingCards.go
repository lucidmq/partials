package pricingCards

import "log"

const DATA_KEY = "pricingCards"
const CONVERTER_FUNCTION_NAME = "cast_to_pricing_cards"

type PricingCard struct {
	Title       string
	Description string
	Price       string
	Features    []string
}

type PricingCards struct {
	PricingCard1 PricingCard
	PricingCard2 PricingCard
	PricingCard3 PricingCard
}

func Cast_to_pricing_cards(mapper map[string]interface{}) PricingCards {
	var v PricingCards
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(PricingCards)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() PricingCards {
	pricingCard := PricingCard{
		Title:       "My Title",
		Description: "My description",
		Price:       "100",
		Features:    []string{"feature1", "feature2", "feature3"},
	}

	return PricingCards{
		PricingCard1: pricingCard,
		PricingCard2: pricingCard,
		PricingCard3: pricingCard,
	}
}
