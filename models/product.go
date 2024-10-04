package models

type ProductID int

type Combination []ChosenStoreProduct

type ChosenStoreProduct struct {
	StoreID   StoreID
	Day       int
	ProductID ProductID
	Price     int
}
