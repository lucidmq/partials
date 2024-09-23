package bizitems

import (
	"log"
)

const DATA_KEY = "bizitems"
const CONVERTER_FUNCTION_NAME = "cast_to_bizitems"

type BizItem struct {
	Title          string
	LogoPath       string
	Link           string
	Description    string
	MonthlyRevenue float32
}

func Cast_to_bizitems(mapper map[string]interface{}) []BizItem {
	var v []BizItem
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.([]BizItem)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
