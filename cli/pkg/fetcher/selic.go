// profitability-cli-go/cli/pkg/fetcher/selic.go
package fetcher

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

var UrlSelic = "https://api.bcb.gov.br/dados/serie/bcdata.sgs.432/dados/ultimos/1?formato=json"

func FetchSelic() (float64, error) {
	response, err := http.Get(UrlSelic)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	if len(data) > 0 {
		selicStr, ok := data[0]["valor"].(string)
		if !ok {
			return 0, fmt.Errorf("valor não é uma string")
		}

		selic, err := strconv.ParseFloat(selicStr, 64)
		if err != nil {
			return 0, err
		}

		return selic, nil
	}

	return 0, fmt.Errorf("nenhum dado retornado")
}

