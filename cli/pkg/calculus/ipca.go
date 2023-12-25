package calculus

import (
	"fmt"
	"errors"
	"profitability/cli/pkg/fetcher"
	"profitability/cli/pkg/util"
)

type ResultIPCA struct {
	ResultIPCA	float64
	TaxRate		float64
}

func IPCA(rate float64, term int) (ResultIPCA, error) {
	if rate <= 0 || term <= 0 {
		return ResultIPCA{}, errors.New("Os valores devem ser maiores que zero.")
	}

	ipca, err := fetcher.FetchIPCA()
	if err != nil {
		return ResultIPCA{}, fmt.Errorf("Falha ao obter o IPCA: %v", err)
	}
    
	taxRate, err := util.Aliquot(term)
	if err != nil {
		return ResultIPCA{}, errors.New("Error: " + err.Error())
	}

	return ResultIPCA{
		ResultIPCA: (rate + ipca / 100) * (1 - taxRate), 
		TaxRate: taxRate * 100,
	}, nil
}
