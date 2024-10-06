package bruteforce_prevalid

import (
	"math"
	"slices"

	"github.com/Sycri/DatZM014-MPD/models"
)

type Solver struct{}

func (*Solver) generateAllValidCombinations(stores *[]models.Store) *[]models.Combination {
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
			// Create combinations with all other elements that have different product IDs
			if !slices.ContainsFunc(combination, func(element models.ChosenStoreProduct) bool {
				return element.ProductID == initialElement.ProductID
			}) {
				combination = append(combination, initialElement)
				newCombinations = append(newCombinations, combination)
			}
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
	allCombinations := s.generateAllValidCombinations(&problem.Stores)
	for _, newCombination := range *allCombinations {
		if valid, newTotalCost, newTotalProductCost, newUsedDayCount := newCombination.CalculateTotalCost(
			&problem.Basket,
		); valid {
			// Check if this is the new best combination
			if solution.TotalCost > newTotalCost {
				solution.TotalCost = newTotalCost
				solution.TotalProductCost = newTotalProductCost
				solution.UsedDayCount = newUsedDayCount
				solution.Combination = newCombination
			}
		}
	}

	return solution
}
