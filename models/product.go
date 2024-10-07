package models

import (
	"math"
	"slices"

	"github.com/Sycri/DatZM014-MPD/constants"
)

type ProductID int

type Combination []ChosenStoreProduct

type ChosenStoreProduct struct {
	StoreID   StoreID
	Day       int
	ProductID ProductID
	Price     int
}

func (c *Combination) CalculateCost(basket *Basket, validate bool) (bool, int64, int64, int) {
	cost := int64(0)
	usedDays := make(map[int]bool, len(*c))

	if validate {
		for _, basketProduct := range basket.Products {
			// Reject combination if a product in the basket is not in this combination
			if !slices.ContainsFunc(*c, func(element ChosenStoreProduct) bool {
				return basketProduct.ID == element.ProductID
			}) {
				return false, int64(math.MaxInt64), -1, -1
			}
		}
	}

	for _, element := range *c {
		basketProductIndex := slices.IndexFunc(basket.Products, func(basketProduct BasketProduct) bool {
			return element.ProductID == basketProduct.ID
		})

		if basketProductIndex == -1 {
			// Should never reach this point if validate is true
			return false, int64(math.MaxInt64), -1, -1
		}

		cost += int64(element.Price * basket.Products[basketProductIndex].Quantity)
		usedDays[element.Day] = true
	}

	productCost := cost
	usedDayCount := len(usedDays)

	// Soft constraint: apply penalty if more days are used than allowed
	if usedDayCount > basket.SoftMaxDays {
		cost += constants.PenaltyPerExtraDay * int64(usedDayCount-basket.SoftMaxDays)
	}

	return true, cost, productCost, usedDayCount
}
