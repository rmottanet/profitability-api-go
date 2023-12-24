package calculus

import (
	"errors"
	"profitability/cli/pkg/util"
)

type ResultPre struct {
	ResultPre	float64
	TaxRate		float64
}

func Pre(rate float64, term int) (ResultPre, error) {
	if rate <= 0 {
		return ResultPre{}, errors.New("A taxa deve ser maior que zero.")
	}
	
	if term <= 0 {
		return ResultPre{}, errors.New("Prazo deve ser maior que zero")
	}	

	taxRate, err := util.Aliquot(term)
	if err != nil {
		return ResultPre{}, errors.New("Error: " + err.Error())
	}

	return ResultPre{
		ResultPre: rate * (1 - taxRate),
		TaxRate: taxRate * 100,
	}, nil
}
