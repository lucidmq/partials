package product

import (
	"log"
)

const DATA_KEY = "product"
const CONVERTER_FUNCTION_NAME = "cast_to_product"

type Details struct {
	Id        string
	Title     string
	KeyPoints []string
}

type Product struct {
	ProductName        string
	ProductDescription string
	ButtonText         string
	Price              int
	ButtonLink         string
	ButtonEnabled      bool
	ImagesLinks        []string
	ProductDetails     []Details
}

func Cast_to_product(mapper map[string]interface{}) Product {
	var v Product
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.(Product)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}
