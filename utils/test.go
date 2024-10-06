package utils

import (
	"slices"
	"testing"

	"github.com/Sycri/DatZM014-MPD/models"
)

func CompareSolve(t *testing.T, expected *models.Solution, actual *models.Solution) {
	if actual.Cost != expected.Cost {
		t.Errorf("Cost expected %d, got %d", expected.Cost, actual.Cost)
	}

	if actual.ProductCost != expected.ProductCost {
		t.Errorf("ProductCost expected %d, got %d", expected.ProductCost, actual.ProductCost)
	}

	if actual.UsedDayCount != expected.UsedDayCount {
		t.Errorf("UsedDayCount expected %d, got %d", expected.UsedDayCount, actual.UsedDayCount)
	}

	if len(actual.Combination) != len(expected.Combination) {
		t.Errorf("Combination length expected %d, got %d", len(expected.Combination), len(actual.Combination))
	}

	slices.SortFunc(actual.Combination, func(a, b models.ChosenStoreProduct) int {
		return int(a.ProductID - b.ProductID)
	})

	slices.SortFunc(expected.Combination, func(a, b models.ChosenStoreProduct) int {
		return int(a.ProductID - b.ProductID)
	})

	for i, element := range actual.Combination {
		if element != expected.Combination[i] {
			t.Errorf("Combination element %d expected %+v, got %+v", i, expected.Combination[i], element)
		}
	}
}
