package fetcher

import (
	"fmt"
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/api"
)

func FetchIPCA() (float64, error) {
	response, err := FetchData(api.UrlIPCA)
	if err != nil {
		return 0, err
	}
	
	ipca, err := util.ParseIPCA(response)
	if err != nil {
		return 0, fmt.Errorf("erro ao processar dados IPCA: %v", err)
	}

	return ipca, nil
}
