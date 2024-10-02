package main

type Solution struct {
	ChosenProducts ChosenProducts
	TotalDays      int
	TotalPrice     int
}

func (s *Solution) CalculateTotalDays() int {
	totalDays := 0

	days := map[int]bool{}

	for _, productStoreInfo := range s.ChosenProducts {
		if _, ok := days[productStoreInfo.Day]; !ok {
			totalDays++
		}

		days[productStoreInfo.Day] = true
	}

	return totalDays
}

func (s *Solution) CalculateTotalPrice(basket *Basket) int {
	totalPrice := 0

	for productID, productStoreInfo := range s.ChosenProducts {
		for _, productOrder := range basket.Products {
			if productOrder.ID != productID {
				continue
			}

			totalPrice += productStoreInfo.Price * productOrder.Quantity
		}
	}

	return totalPrice
}

func (p *Problem) BruteSolve() *Solution {
	s := &Solution{
		ChosenProducts: make(ChosenProducts, len(p.Basket.Products)),
	}

	for _, basketProduct := range p.Basket.Products {
		for _, store := range p.Stores {
			for day, dayProducts := range store.DayOfferings {
				for _, storeProduct := range dayProducts {
					if storeProduct.ID != basketProduct.ID {
						continue
					}

					chosenProduct, ok := s.ChosenProducts[basketProduct.ID]
					if !ok || storeProduct.Price < chosenProduct.Price {
						s.ChosenProducts[basketProduct.ID] = ProductStoreInfo{
							StoreID: store.ID,
							Day:     day,
							Price:   storeProduct.Price,
						}
					}
				}
			}
		}
	}

	s.TotalDays = s.CalculateTotalDays()
	s.TotalPrice = s.CalculateTotalPrice(&p.Basket)

	return s
}
