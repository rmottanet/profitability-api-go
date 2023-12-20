package util

import (
	"errors"
)

type Tribute struct {
	UpperBound int
	TaxRate    float64
}

func Aliquot(term int) (float64, error) {
	taxRanges := []Tribute{
		{UpperBound: 180, TaxRate: 0.225},
		{UpperBound: 360, TaxRate: 0.20},
		{UpperBound: 720, TaxRate: 0.175},
		{UpperBound: int(^uint(0) >> 1), TaxRate: 0.15},
	}

	for _, tribute := range taxRanges {
		if term <= tribute.UpperBound {
			return tribute.TaxRate, nil
		}
	}

	return 0.00, errors.New("No tax rate found for the given term")
}

