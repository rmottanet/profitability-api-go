package calculus

import (
	"fmt"
	"errors"
	"profitability/cli/pkg/fetcher"
)

type ResultProp struct {
	ResultProp	float64
	Selic		float64
}

func Prop(rate float64) (ResultProp, error) {
	if rate <= 0 {
		return ResultProp{}, errors.New("A taxa deve ser maior que zero.")
	}
	
	selic, err := fetcher.FetchSelic()
	if err != nil {
		return ResultProp{}, fmt.Errorf("Falha ao obter a taxa SELIC: %v", err)
	}

	return ResultProp{
		ResultProp:  rate * selic / 100,
		Selic:  selic,
	}, nil
}
