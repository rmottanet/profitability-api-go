package calculus

import (
	"fmt"
	"math"
	"errors"
	"profitability/cli/pkg/fetcher"
	"profitability/cli/pkg/util"
)

type ResultPos struct {
	ResultPos	float64
	TaxRate		float64
}

func Pos(rate float64, term int) (ResultPos, error) {
	if rate <= 0 {
		return ResultPos{}, errors.New("A taxa deve ser maior que zero.")
	}
	
	if term <= 0 {
		return ResultPos{}, errors.New("Prazo deve ser maior que zero")
	}	

	selic, err := fetcher.FetchSelic()
	if err != nil {
		return ResultPos{}, fmt.Errorf("Falha ao obter a taxa SELIC: %v", err)
	}

	taxRate, err := util.Aliquot(term)
	if err != nil {
		return ResultPos{}, errors.New("Error: " + err.Error())
	}

	result := (rate * selic / 100) * (1 - taxRate)
	result = math.Round(result*100) / 100

	return ResultPos{
		ResultPos: result, 
		TaxRate: taxRate * 100,
	}, nil
}
