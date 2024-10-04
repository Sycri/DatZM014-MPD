package models

import (
	"encoding/json"
	"fmt"
)

type StoreID int
type DayOfferings map[int][]StoreProduct

type Store struct {
	ID           StoreID
	Name         string
	DayOfferings DayOfferings
}

type StoreProduct struct {
	ID    ProductID
	Price int
}

// Tranforms JSON array into Go map instead of Go slice
func (df *DayOfferings) UnmarshalJSON(data []byte) error {
	var tmp []struct {
		Day      int
		Products []StoreProduct
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*df = make(DayOfferings, len(tmp))
	for _, obj := range tmp {
		if _, ok := (*df)[obj.Day]; ok {
			return fmt.Errorf("day %d already exists", obj.Day)
		}

		(*df)[obj.Day] = obj.Products
	}

	return nil
}
