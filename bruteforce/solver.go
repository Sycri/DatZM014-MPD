package bruteforce

import (
	"math"

	"github.com/Sycri/DatZM014-MPD/models"
)

type Solver struct{}

func (*Solver) calculateUsedDayCount(combination *models.Combination) int {
	days := make(map[int]bool, len(*combination))

	for _, storeDayProduct := range *combination {
		days[storeDayProduct.Day] = true
	}

	return len(days)
}

func (*Solver) calculateTotalProductCost(
	combination *models.Combination,
	basketProducts *[]models.BasketProduct,
) int64 {
	totalProductCost := int64(0)

	for _, storeDayProduct := range *combination {
		for _, basketProduct := range *basketProducts {
			if basketProduct.ID == storeDayProduct.ProductID {
				totalProductCost += int64(storeDayProduct.Price * basketProduct.Quantity)
				break
			}
		}
	}

	return totalProductCost
}

func (*Solver) generateAllCombinations(stores *[]models.Store) *[]models.Combination {
	if len(*stores) == 0 {
		return &[]models.Combination{}
	}

	// Flatten store products in each day to a single array
	initialElements := []models.ChosenStoreProduct{}
	for _, store := range *stores {
		for day, products := range store.DayOfferings {
			for _, product := range products {
				initialElements = append(initialElements, models.ChosenStoreProduct{
					StoreID:   store.ID,
					Day:       day,
					ProductID: product.ID,
					Price:     product.Price,
				})
			}
		}
	}

	combinations := []models.Combination{}

	// Initialize each combination with first element
	for _, initialElement := range initialElements {
		combinations = append(combinations, models.Combination{initialElement})
	}

	// Generate all combinations
	for _, initialElement := range initialElements {
		newCombinations := []models.Combination{}

		for _, combination := range combinations {
			// Create combinations with all other elements
			combination = append(combination, initialElement)
			newCombinations = append(newCombinations, combination)
		}

		combinations = append(combinations, newCombinations...)
	}

	return &combinations
}

func (s *Solver) Solve(problem *models.Problem) *models.Solution {
	solution := &models.Solution{
		TotalCost: int64(math.MaxInt64),
	}

	// Iterate over all possible combinations
	allCombinations := s.generateAllCombinations(&problem.Stores)
	for _, newCombination := range *allCombinations {
		if valid, newTotalCost := newCombination.CalculateTotalCost(&problem.Basket); valid {
			// Check if this is the new best combination
			if solution.TotalCost > newTotalCost {
				solution.TotalCost = newTotalCost
				solution.Combination = newCombination
			}
		}
	}

	solution.UsedDayCount = s.calculateUsedDayCount(&solution.Combination)
	solution.TotalProductCost = s.calculateTotalProductCost(&solution.Combination, &problem.Basket.Products)

	return solution
}
