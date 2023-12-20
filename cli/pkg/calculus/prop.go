package calculus

import (
	"profitability/cli/pkg/fetcher"
)

type ResultProp struct {
	ResultProp float64
	Selic  float64
}

func Prop(rate float64) (ResultProp, error) {
	selic, err := fetcher.FetchSelic()
	if err != nil {
		return ResultProp{}, err
	}

	result := rate * selic / 100
	return ResultProp{
		ResultProp: result,
		Selic:  selic,
	}, nil
}
