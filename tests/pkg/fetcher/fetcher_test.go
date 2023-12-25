package fetcher_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"profitability/cli/pkg/fetcher"
)

func TestFetchData(t *testing.T) {
	// Criar um servidor de teste simulado
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simular uma resposta JSON
		responseJSON := `[{"key1": "value1"}, {"key2": "value2"}]`
		w.Write([]byte(responseJSON))
	}))

	defer server.Close()

	// Testar a função FetchData com a URL do servidor de teste
	data, err := fetcher.FetchData(server.URL)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	// Verificar o tamanho do slice retornado
	if len(data) != 2 {
		t.Errorf("Tamanho do slice inesperado. Esperado 2, mas obtido %d", len(data))
	}

	// Verificar o conteúdo do slice retornado
	expectedData := map[string]interface{}{"key1": "value1"}
	if data[0]["key1"] != expectedData["key1"] {
		t.Errorf("Valor inesperado. Esperado %v, mas obtido %v", expectedData, data[0])
	}

	expectedData = map[string]interface{}{"key2": "value2"}
	if data[1]["key2"] != expectedData["key2"] {
		t.Errorf("Valor inesperado. Esperado %v, mas obtido %v", expectedData, data[1])
	}
}
