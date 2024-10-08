package utils

import (
	"math/big"
)

func PowerSetFunc[S ~[]E, E any](elements []E, skipEmptySet bool, f func(S)) {
	if !skipEmptySet && elements != nil {
		f(S{})
	}

	elementCount := len(elements)
	if elementCount > 63 {
		bigPowerSetFunc(elements, f)
		return
	}

	powerSetSize := uint64Pow(2, elementCount)

	// Starting with 1 because 0 is the empty set
	for i := uint64(1); i < powerSetSize; i++ {
		var newSet S

		for j, elem := range elements {
			if (i>>j)&1 > 0 {
				newSet = append(newSet, elem)
			}
		}

		f(newSet)
	}
}

func bigPowerSetFunc[S ~[]E, E any](elements []E, f func(S)) {
	base, exponent := big.NewInt(2), big.NewInt(int64(len(elements)))
	powerSetSize := base.Exp(base, exponent, nil)

	// Starting with 1 because 0 is the empty set
	for i := big.NewInt(1); i.Cmp(powerSetSize) < 0; i.Add(i, big.NewInt(1)) {
		var newSet S

		for j, elem := range elements {
			if i.Bit(j) > 0 {
				newSet = append(newSet, elem)
			}
		}

		f(newSet)
	}
}

func uint64Pow(n uint64, m int) uint64 {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
