package calculus

import (
	"math"
	"errors"
	"profitability/cli/pkg/util"
)

type ResultPre struct {
	ResultPre   float64
	TaxRate  float64
}

func Pre(rate float64, term int) (ResultPre, error) {
	if term <= 0 {
		return ResultPre{}, errors.New("Prazo deve ser maior que zero")
	}	

	taxRate, err := util.Aliquot(term)
	if err != nil {
		return ResultPre{}, errors.New("Error: " + err.Error())
	}

	result := rate * (1 - taxRate)
	result = math.Round(result*100) / 100

	return ResultPre{ResultPre: result, TaxRate: taxRate * 100}, nil
}

