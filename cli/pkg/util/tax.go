package util

import (
	"errors"
)

type Tribute struct {
	UpperBound int
	TaxRate    float64
}

const (
    initialTaxRate = 0.00
    lowerBound180  = 0
    upperBound180  = 180
    upperBound360  = 360
    upperBound720  = 720
)

func Aliquot(term int) (float64, error) {
	if term < 0 {
	    return initialTaxRate, errors.New("O prazo não pode ser negativo")
	}

	taxRanges := []Tribute{
	    {UpperBound: upperBound180, TaxRate: 0.225},
	    {UpperBound: upperBound360, TaxRate: 0.20},
	    {UpperBound: upperBound720, TaxRate: 0.175},
	    {UpperBound: int(^uint(0) >> 1), TaxRate: 0.15},
	}
	
	for _, tribute := range taxRanges {
		if term <= tribute.UpperBound {
			return tribute.TaxRate, nil
		}
	}

	return initialTaxRate, errors.New("No tax rate found for the given term")
}

