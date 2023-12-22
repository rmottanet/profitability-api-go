// tests/fetcher_test.go
package fetcher_test

import (
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"profitability/cli/pkg/fetcher"
)

func TestFetchIPCA(t *testing.T) {
	// Criar um servidor de teste para simular a resposta da API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Resposta simulada da API
		responseJSON := `[{"valor": "3.50"}, {"valor": "3.80"}, {"valor": "4.10"}, {"valor": "-4.30"}, {"valor": "4.50"}, {"valor": "4.81"}, {"valor": "5.10"}, {"valor": "5.40"}, {"valor": "5.72"}, {"valor": "-0.02"}, {"valor": "6.20"}, {"valor": "6.50"}]`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJSON))
	}))
	defer server.Close()

	// Substituir a URL real pela URL do servidor de teste
	oldURL := fetcher.UrlIPCA
	fetcher.UrlIPCA = server.URL
	defer func() { fetcher.UrlIPCA = oldURL }()

	// Chamar a função FetchIPCA e verificar o resultado
	ipca, err := fetcher.FetchIPCA()
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	expectedIPCA := 45.31
	if math.Abs(expectedIPCA - ipca) > 0.0001 {
		t.Errorf("Valor do IPCA esperado: %f, Obtido: %f", expectedIPCA, ipca)
	}
}
