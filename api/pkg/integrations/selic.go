// Fetcher SELIC
package integrations

import (
    "fmt"
    "log"
    "time"
    "encoding/json"

    "profitability/api/pkg/cache"
    "profitability/api/pkg/util"
)


type SELICResponse struct {
    Data  string  `json:"data"`
    Valor float64 `json:"valor"`
}

func GetSELICData(cache cache.Cache) (*SELICResponse, error) {
    apiURL := GetUrl("selic")

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

    selicData := parsedData[0]

    selic := SELICResponse{
        Data:  selicData["data"].(string),
    }

    // processa dados usando a função util.parseselic
    selic.Valor, err = util.ParseSelic(parsedData)
    if err != nil {
        return nil, fmt.Errorf("error parsing SELIC value: %v", err)
    }
	
	// cache dados
    var expiration = 12 * time.Hour

    cache.Set("SELIC", selic.Valor, expiration)
	log.Printf("SELIC cached %f", selic.Valor)

    return &selic, nil
}
