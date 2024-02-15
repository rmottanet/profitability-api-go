package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

// FetchData faz uma requisição HTTP para a URL fornecida e decodifica a resposta JSON.
func FetchData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição HTTP: %v", err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a resposta HTTP: %v", err)
	}

	return data, nil
}
