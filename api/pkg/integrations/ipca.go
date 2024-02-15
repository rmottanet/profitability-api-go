// Fetcher IPCA
package integrations

import (
    "fmt"
    "log"
    "time"
    "encoding/json"

    "profitability/api/pkg/cache"
    "profitability/api/pkg/util"
)

type IPCAResponse struct {
    Data  string  `json:"data"`
    Valor float64 `json:"valor"`
}

func GetIPCAData(cache cache.Cache) (*IPCAResponse, error) {
    apiURL := GetUrl("ipca")

    data, err := util.FetchData(apiURL)
    if err != nil {
        return nil, fmt.Errorf("error getting data from API: %v", err)
    }

    // processa resposta
    var parsedData []map[string]interface{}
    err = json.Unmarshal(data, &parsedData)
    if err != nil {
        return nil, fmt.Errorf("error decoding JSON response: %v", err)
    }

    if len(parsedData) == 0 {
        return nil, fmt.Errorf("empty response data from API")
    }

    ipcaData := parsedData[0]

    ipca := IPCAResponse{
        Data:  ipcaData["data"].(string),
    }

    // acumula os dados usando a função util.parseipca
    ipca.Valor, err = util.ParseIPCA(parsedData)
    if err != nil {
        return nil, fmt.Errorf("error parsing IPCA value: %v", err)
    }

	// cache dados
    var expiration = 12 * time.Hour

    cache.Set("IPCA", ipca.Valor, expiration)
    log.Printf("IPCA cached %f", ipca.Valor)

    return &ipca, nil
}
