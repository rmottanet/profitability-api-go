package fetcher

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// FetchData faz uma requisição HTTP para a URL fornecida e decodifica a resposta JSON.
func FetchData(url string) ([]map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição HTTP: %v", err)
	}
	defer response.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	return data, nil
}
