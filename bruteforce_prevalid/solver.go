package bruteforce_prevalid

import (
	"math"
	"slices"

	"github.com/Sycri/DatZM014-MPD/models"
)

type Solver struct{}

func (*Solver) getFlattenedElements(stores *[]models.Store) *[]models.ChosenStoreProduct {
	flattenedElements := []models.ChosenStoreProduct{}

	// Flatten store products in each day to a single array
	for _, store := range *stores {
		for day, products := range store.DayOfferings {
			for _, product := range products {
				flattenedElements = append(flattenedElements, models.ChosenStoreProduct{
					StoreID:   store.ID,
					Day:       day,
					ProductID: product.ID,
					Price:     product.Price,
				})
			}
		}
	}

	return &flattenedElements
}

func (s *Solver) generateAllValidCombinations(stores *[]models.Store) *[]models.Combination {
	if len(*stores) == 0 {
		return &[]models.Combination{}
	}

	initialElements := s.getFlattenedElements(stores)

	combinations := []models.Combination{}

	// Initialize each combination with first element
	for _, initialElement := range *initialElements {
		combinations = append(combinations, models.Combination{initialElement})
	}

	// Generate all combinations
	for _, initialElement := range *initialElements {
		newCombinations := []models.Combination{}

		for _, combination := range combinations {
			// Create combinations with all other elements that have different product IDs
			if !slices.ContainsFunc(combination, func(element models.ChosenStoreProduct) bool {
				return element.ProductID == initialElement.ProductID
			}) {
				newCombination := append(make([]models.ChosenStoreProduct, 0, len(combination)+1), combination...)
				newCombination = append(newCombination, initialElement)
				newCombinations = append(newCombinations, newCombination)
			}
		}

		combinations = append(combinations, newCombinations...)
	}

	return &combinations
}

func (s *Solver) Solve(problem *models.Problem) *models.Solution {
	solution := &models.Solution{
		Cost: int64(math.MaxInt64),
	}

	// Iterate over all valid combinations
	combinations := s.generateAllValidCombinations(&problem.Stores)
	for _, newCombination := range *combinations {
		if valid, newCost, newProductCost, newUsedDayCount := newCombination.CalculateCost(
			&problem.Basket,
		); valid {
			// Check if this is the new best combination
			if solution.Cost > newCost {
				solution.Cost = newCost
				solution.ProductCost = newProductCost
				solution.UsedDayCount = newUsedDayCount
				solution.Combination = newCombination
			}
		}
	}

	return solution
}
