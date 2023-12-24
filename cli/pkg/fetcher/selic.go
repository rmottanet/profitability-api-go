package fetcher

import (
	"fmt"
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/api"
)

func FetchSelic() (float64, error) {
	response, err := FetchData(api.UrlSelic)
	if err != nil {
		return 0, err
	}

	selic, err := util.ParseSelic(response)
	if err != nil {
		return 0, fmt.Errorf("erro ao processar dados IPCA: %v", err)
	}
	
	return selic, nil
}
