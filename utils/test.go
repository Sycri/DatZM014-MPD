package utils

import (
	"slices"
	"testing"

	"github.com/Sycri/DatZM014-MPD/models"
)

func CompareSolve(t *testing.T, expected *models.Solution, actual *models.Solution) {
	if actual.TotalCost != expected.TotalCost {
		t.Errorf("TotalCost expected %d, got %d", expected.TotalCost, actual.TotalCost)
	}

	if actual.TotalProductCost != expected.TotalProductCost {
		t.Errorf("TotalProductCost expected %d, got %d", expected.TotalProductCost, actual.TotalProductCost)
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
