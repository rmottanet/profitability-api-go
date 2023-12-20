package fetcher

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

var UrlIPCA = "https://api.bcb.gov.br/dados/serie/bcdata.sgs.433/dados/ultimos/12?formato=json"

func FetchIPCA() (float64, error) {
	//url := "https://api.bcb.gov.br/dados/serie/bcdata.sgs.433/dados/ultimos/12?formato=json"
	response, err := http.Get(UrlIPCA)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	var ipca float64
	for i := 0; i < len(data); i++ {
		valorStr, ok := data[i]["valor"].(string)
		if !ok {
			return 0, fmt.Errorf("valor não é uma string")
		}

		valor, err := strconv.ParseFloat(valorStr, 64)
		if err != nil {
			return 0, err
		}

		ipca += valor
	}

//	ipcaFormatted := fmt.Sprintf("%.2f", ipca)

	return ipca, nil
}
