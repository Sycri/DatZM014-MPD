package simulated_annealing

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/Sycri/DatZM014-MPD/models"
)

const (
	initialTemperature = 1000.0
	coolingRate        = 0.99
	maxIterations      = 10000
)

type Solver struct{}

func (*Solver) getStoresPerProducts(stores *[]models.Store) *map[models.ProductID][]models.ChosenStoreProduct {
	storesPerProducts := make(map[models.ProductID][]models.ChosenStoreProduct)

	// Map stores to products
	for _, store := range *stores {
		for day, products := range store.DayOfferings {
			for _, product := range products {
				productStores, ok := storesPerProducts[product.ID]
				if !ok {
					productStores = []models.ChosenStoreProduct{}
				}

				storesPerProducts[product.ID] = append(productStores, models.ChosenStoreProduct{
					StoreID:   store.ID,
					Day:       day,
					ProductID: product.ID,
					Price:     product.Price,
				})
			}
		}
	}

	return &storesPerProducts
}

func (*Solver) getInitialRandomCombination(
	basketProducts *[]models.BasketProduct,
	storesPerProducts *map[models.ProductID][]models.ChosenStoreProduct,
) *models.Combination {
	combination := make(models.Combination, len(*basketProducts))

	for i, basketProduct := range *basketProducts {
		productStores, ok := (*storesPerProducts)[basketProduct.ID]
		if !ok {
			panic(fmt.Errorf("product %d not found among stores", basketProduct.ID))
		}

		// Find random store + day for product
		randomProductStoreIndex := rand.Intn(len(productStores))
		combination[i] = productStores[randomProductStoreIndex]
	}

	return &combination
}

func (*Solver) mutateCombination(
	combination *models.Combination,
	storesPerProducts *map[models.ProductID][]models.ChosenStoreProduct,
) *models.Combination {
	newCombination := make(models.Combination, len(*combination))
	copy(newCombination, *combination)

	// Randomly change the store + day for one product
	i := rand.Intn(len(newCombination))

	productStores, ok := (*storesPerProducts)[(*combination)[i].ProductID]
	if !ok {
		panic(fmt.Errorf("product %d not found among stores", (*combination)[i].ProductID))
	}

	// Find random store + day for product
	randomProductStoreIndex := rand.Intn(len(productStores))
	newCombination[i] = productStores[randomProductStoreIndex]

	return &newCombination
}

func (*Solver) acceptNewSolution(currentCost int64, newCost int64, temperature float64) bool {
	if newCost < currentCost {
		return true
	}

	acceptanceProbability := math.Exp(float64(currentCost-newCost) / temperature)
	return rand.Float64() < acceptanceProbability
}

func (s *Solver) Solve(problem *models.Problem) *models.Solution {
	storesPerProducts := s.getStoresPerProducts(&problem.Stores)

	// Generate initial solution
	bestSolution := &models.Solution{
		Combination: *s.getInitialRandomCombination(&problem.Basket.Products, storesPerProducts),
	}

	// Calculate initial solution cost & other stats
	_, bestSolution.Cost, bestSolution.ProductCost, bestSolution.UsedDayCount = bestSolution.Combination.CalculateCost(
		&problem.Basket, false,
	)

	currentSolution := bestSolution
	temperature := initialTemperature

	for i := 0; i < maxIterations; i++ {
		// Mutate the current solution to get a new solution
		newSolution := &models.Solution{
			Combination: *s.mutateCombination(&currentSolution.Combination, storesPerProducts),
		}
		_, newSolution.Cost, newSolution.ProductCost, newSolution.UsedDayCount = newSolution.Combination.CalculateCost(
			&problem.Basket, false,
		)

		// If the new solution is better or te, accept it
		if s.acceptNewSolution(currentSolution.Cost, newSolution.Cost, temperature) {
			currentSolution = newSolution

			// Update the best solution if needed
			if currentSolution.Cost < bestSolution.Cost {
				bestSolution = currentSolution
			}
		}

		// Cool down the temperature
		temperature *= coolingRate
	}

	return bestSolution
}
