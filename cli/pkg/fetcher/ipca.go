package fetcher

import (
	"fmt"
	"net/http"
	"encoding/json"
	"profitability/cli/pkg/util"
)

var UrlIPCA = "https://api.bcb.gov.br/dados/serie/bcdata.sgs.433/dados/ultimos/12?formato=json"

func FetchIPCA() (float64, error) {
	response, err := http.Get(UrlIPCA)
	if err != nil {
		return 0, fmt.Errorf("erro ao fazer a requisição HTTP: %v", err)
	}
	defer response.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	if len(data) == 0 {
		return 0, fmt.Errorf("nenhum dado retornado")
	}
	
	ipca, err := util.ParseIPCA(data)
	if err != nil {
		return 0, fmt.Errorf("erro ao processar dados IPCA: %v", err)
	}

	return ipca, nil
}
