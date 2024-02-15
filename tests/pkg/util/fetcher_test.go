package util_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"profitability/api/pkg/util"
)

func TestFetchData(t *testing.T) {
	// Criando um servidor de teste
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Escrevendo uma resposta de teste
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer testServer.Close()

	// Testando a função FetchData com o servidor de teste
	url := testServer.URL
	data, err := util.FetchData(url)
	if err != nil {
		t.Fatalf("Erro inesperado ao buscar dados: %v", err)
	}

	// Verificando se os dados retornados não estão vazios
	if len(data) == 0 {
		t.Error("FetchData retornou dados vazios")
	}

	// Verificando se os dados retornados são os esperados
	expected := []byte(`{"key": "value"}`)
	if string(data) != string(expected) {
		t.Errorf("FetchData retornou dados diferentes do esperado. Esperado: %s, Obtido: %s", expected, data)
	}
}
