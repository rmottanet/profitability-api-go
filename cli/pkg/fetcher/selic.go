package fetcher

import (
	"fmt"
	"net/http"
	"encoding/json"
	"profitability/cli/pkg/util"
)

var UrlSelic = "https://api.bcb.gov.br/dados/serie/bcdata.sgs.432/dados/ultimos/1?formato=json"

func FetchSelic() (float64, error) {
	response, err := http.Get(UrlSelic)
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
	
	selicStr, ok := data[0]["valor"].(string)
	if !ok {
		return 0, fmt.Errorf("valor não é uma string")
	}

	selic, err := util.ParseFloat(selicStr)
	if err != nil {
		return 0, fmt.Errorf("erro ao processar dados SELIC: %v", err)
	}
	
	return selic, nil
}
