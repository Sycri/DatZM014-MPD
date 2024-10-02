package main

import (
	"encoding/json"
)

type ProductID int

type Product struct {
	ID   ProductID
	Name string
}

type StoreProduct struct {
	ID    ProductID
	Price int
}

type ProductOrder struct {
	ID       ProductID
	Quantity int
}

type ChosenProducts map[ProductID]ProductStoreInfo

type ProductStoreInfo struct {
	StoreID StoreID
	Day     int
	Price   int
}

// Transforms Go map into JSON array without keys
func (cp *ChosenProducts) MarshalJSON() ([]byte, error) {
	type UnmappedChosenProduct struct {
		ProductID
		ProductStoreInfo
	}

	tmp := make([]UnmappedChosenProduct, 0, len(*cp))
	for productID, productStoreInfo := range *cp {
		tmp = append(tmp, UnmappedChosenProduct{
			ProductID:        productID,
			ProductStoreInfo: productStoreInfo,
		})
	}

	return json.Marshal(tmp)
}
