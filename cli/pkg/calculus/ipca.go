package calculus

import (
	"math"
	"errors"
	"profitability/cli/pkg/fetcher"
	"profitability/cli/pkg/util"
)

type ResultIPCA struct {
	ResultIPCA   float64
	TaxRate  float64
}

func IPCA(rate float64, term int) (ResultIPCA, error) {
	if term <= 0 {
		return ResultIPCA{}, errors.New("Prazo deve ser maior que zero")
	}

	ipca, err := fetcher.FetchIPCA()
	if err != nil {
		return ResultIPCA{}, err
	}
    
	taxRate, err := util.Aliquot(term)
	if err != nil {
		return ResultIPCA{}, errors.New("Error: " + err.Error())
	}

	result := (rate + ipca / 100) * (1 - taxRate)
	result = math.Round(result*100) / 100

	return ResultIPCA{ResultIPCA: result, TaxRate: taxRate * 100}, nil
}
