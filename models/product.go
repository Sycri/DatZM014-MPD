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

func (c *Combination) CalculateCost(basket *Basket) (bool, int64, int64, int) {
	cost := int64(0)
	usedDays := make(map[int]bool, len(*c))

	for _, basketProduct := range basket.Products {
		elementIndex := slices.IndexFunc(*c, func(element ChosenStoreProduct) bool {
			return element.ProductID == basketProduct.ID
		})

		// Reject combination if a product in the basket is not in this combination
		if elementIndex == -1 {
			return false, int64(math.MaxInt64), -1, -1
		}

		cost += int64((*c)[elementIndex].Price * basketProduct.Quantity)
		usedDays[(*c)[elementIndex].Day] = true
	}

	productCost := cost
	usedDayCount := len(usedDays)

	// Soft constraint: apply penalty if more days are used than allowed
	if usedDayCount > basket.SoftMaxDays {
		cost += constants.PenaltyPerExtraDay * int64(usedDayCount-basket.SoftMaxDays)
	}

	return true, cost, productCost, usedDayCount
}
