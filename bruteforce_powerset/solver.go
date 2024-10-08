package bruteforce_powerset

import (
	"math"

	"github.com/Sycri/DatZM014-MPD/models"
	"github.com/Sycri/DatZM014-MPD/utils"
)

type Solver struct{}

func (*Solver) getFlattenedElements(stores *[]models.Store) []*models.ChosenStoreProduct {
	flattenedElements := []*models.ChosenStoreProduct{}

	// Flatten store products in each day to a single array
	for _, store := range *stores {
		for day, products := range store.DayOfferings {
			for _, product := range products {
				flattenedElements = append(flattenedElements, &models.ChosenStoreProduct{
					StoreID:   store.ID,
					Day:       day,
					ProductID: product.ID,
					Price:     product.Price,
				})
			}
		}
	}

	return flattenedElements
}

func (s *Solver) Solve(problem *models.Problem) *models.Solution {
	solution := &models.Solution{
		Cost: int64(math.MaxInt64),
	}

	elements := s.getFlattenedElements(&problem.Stores)

	// Iterate over all possible combinations
	utils.PowerSetFunc(elements, true, func(newCombination models.Combination) {
		if valid, newCost, newProductCost, newUsedDayCount := newCombination.CalculateCost(
			&problem.Basket, true,
		); valid {
			// Check if this is the new best combination
			if solution.Cost > newCost {
				solution.Cost = newCost
				solution.ProductCost = newProductCost
				solution.UsedDayCount = newUsedDayCount
				solution.Combination = newCombination
			}
		}
	})

	return solution
}
