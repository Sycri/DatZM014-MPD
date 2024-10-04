package models

type Basket struct {
	Products    []BasketProduct
	SoftMaxDays int
}

type BasketProduct struct {
	ID       ProductID
	Quantity int
}
