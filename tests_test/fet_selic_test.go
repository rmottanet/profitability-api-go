// tests/fet_selic_test.go
package fetcher_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"profitability/cli/pkg/fetcher"
)


func TestFetchSelic(t *testing.T) {
	// Criar um servidor de teste para simular a resposta da API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Resposta simulada da API
		responseJSON := `[{"valor": "5.25"}]`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJSON))
	}))
	defer server.Close()

	// Substituir a URL real pela URL do servidor de teste
	oldURL := fetcher.UrlSelic
	fetcher.UrlSelic = server.URL
	defer func() { fetcher.UrlSelic = oldURL }()

	// Chamar a função FetchSelic e verificar o resultado
	selic, err := fetcher.FetchSelic()
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	expectedSelic := 5.25
	if selic != expectedSelic {
		t.Errorf("Valor da SELIC esperado: %f, Obtido: %f", expectedSelic, selic)
	}
}
