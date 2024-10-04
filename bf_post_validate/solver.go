package bf_post_validate

import (
	"math"
	"slices"

	"github.com/Sycri/DatZM014-MPD/models"
)

const (
	penaltyPerExtraDay = 100 // Cost per each extra day
)

type Solver struct{}

func (_ *Solver) calculateUsedDayCount(combination *models.Combination) int {
	days := make(map[int]bool, len(*combination))

	for _, storeDayProduct := range *combination {
		days[storeDayProduct.Day] = true
	}

	return len(days)
}

func (_ *Solver) calculateTotalProductCost(
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

func (_ *Solver) generateAllCombinations(stores *[]models.Store) *[]models.Combination {
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

func (_ *Solver) calculateTotalCost(
	basket *models.Basket,
	combination *models.Combination,
) (bool, int64) {
	totalCost := int64(0)
	usedDays := make(map[int]bool, len(*combination))

	for _, basketProduct := range basket.Products {
		elementIndex := slices.IndexFunc(*combination, func(element models.ChosenStoreProduct) bool {
			return element.ProductID == basketProduct.ID
		})

		// Reject combination if a product in the basket is not in this combination
		if elementIndex == -1 {
			return false, int64(math.MaxInt64)
		}

		totalCost += int64((*combination)[elementIndex].Price * basketProduct.Quantity)
		usedDays[(*combination)[elementIndex].Day] = true
	}

	usedDayCount := len(usedDays)

	// Soft constraint: apply penalty if more days are used than allowed
	if usedDayCount > basket.SoftMaxDays {
		totalCost += penaltyPerExtraDay * int64(usedDayCount-basket.SoftMaxDays)
	}

	return true, totalCost
}

func (s *Solver) Solve(problem *models.Problem) *models.Solution {
	solution := &models.Solution{
		TotalCost: int64(math.MaxInt64),
	}

	// Iterate over all possible combinations
	allCombinations := s.generateAllCombinations(&problem.Stores)
	for _, newCombination := range *allCombinations {
		if valid, newTotalCost := s.calculateTotalCost(&problem.Basket, &newCombination); valid {
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
