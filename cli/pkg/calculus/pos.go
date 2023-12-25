package calculus

import (
	"fmt"
	"errors"
	"profitability/cli/pkg/fetcher"
	"profitability/cli/pkg/util"
)

type ResultPos struct {
	ResultPos	float64
	TaxRate		float64
}

func Pos(rate float64, term int) (ResultPos, error) {
	if rate <= 0 || term <= 0 {
		return ResultPre{}, errors.New("Os valores devem ser maiores que zero.")
	}	

	selic, err := fetcher.FetchSelic()
	if err != nil {
		return ResultPos{}, fmt.Errorf("Falha ao obter a taxa SELIC: %v", err)
	}

	taxRate, err := util.Aliquot(term)
	if err != nil {
		return ResultPos{}, errors.New("Error: " + err.Error())
	}

	return ResultPos{
		ResultPos: (rate * selic / 100) * (1 - taxRate), 
		TaxRate: taxRate * 100,
	}, nil
}
