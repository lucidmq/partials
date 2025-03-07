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
	Price              string
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

func NewDummy() Product {
	productDetails := []Details{
		{
			Id:    "nakljd",
			Title: "Features",
			KeyPoints: []string{
				"Polycarbonate frame provides structure",
				"Soft-touch TPU surface",
				"Easily insert/remove your AirTag with snap-in AirTag mount geometry",
				"AirTag not included",
				"The AirTag is as thick as 10 credit cards",
			},
		},
		{
			Id:    "welr",
			Title: "Shipping",
			KeyPoints: []string{
				"Free shipping on orders over $300",
				"International shipping available",
				"Expedited shipping options",
				"Signature required upon delivery",
			},
		},
		{
			Id:    "lajdklfa",
			Title: "Returns",
			KeyPoints: []string{
				"Easy return requests",
				"Pre-paid shipping label included",
				"10% restocking fee for returns",
				"60 day return window",
			},
		},
	}

	productData := Product{
		ProductName:        "AirTag Wallet Card",
		ProductDescription: "The Card for AirTag is a sleek, credit card-shaped holder designed for easy wallet tracking with the Find My app. The Apple AirTag is about the thickness of 10 regular credit cards, and this holder keeps things slim. It won’t add bulk but needs a deep pocket to prevent wallet stretching or bending your cards. Not ideal for card-only or hard shell wallets, but perfect for those who want convenience without the extra bulk!",
		ButtonText:         "Buy Now",
		Price:              "$10.00",
		ButtonEnabled:      true,
		ButtonLink:         "/checkout",
		ImagesLinks:        []string{"/static/images/product_images/airtag-card-w-specs.avif", "/static/images/product_images/airtag-card-w-wallet.avif", "/static/images/product_images/airtag-card-with-description.avif", "/static/images/product_images/black-airtag-card.avif"},
		ProductDetails:     productDetails,
	}
	return productData
}
