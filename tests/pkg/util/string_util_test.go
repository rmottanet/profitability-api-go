package util_test

import (
	"testing"
	"profitability/cli/pkg/util"
)

func TestParseFloat(t *testing.T) {
	// Teste com valor válido
	result, err := util.ParseFloat("10.5")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	expectedResult := 10.5
	if result != expectedResult {
		t.Errorf("Resultado inesperado. Esperado: %f, Obtido: %f", expectedResult, result)
	}

	// Teste com valor inválido
	_, err = util.ParseFloat("abc")
	if err == nil {
		t.Error("Esperava um erro para string inválida, mas nenhum foi retornado.")
	}
}

func TestParseIPCA(t *testing.T) {
	// Teste com dados válidos
	data := []map[string]interface{}{
		{"valor": "2.5"},
		{"valor": "3.0"},
		{"valor": "1.8"},
	}

	result, err := util.ParseIPCA(data)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	expectedResult := 7.3
	if result != expectedResult {
		t.Errorf("Resultado inesperado. Esperado: %f, Obtido: %f", expectedResult, result)
	}

	// Teste com valor inválido
	_, err = util.ParseIPCA([]map[string]interface{}{{"valor": "abc"}})
	if err == nil {
		t.Error("Esperava um erro para string inválida, mas nenhum foi retornado.")
	}

	// Teste com dados vazios
	_, err = util.ParseIPCA([]map[string]interface{}{})
	if err == nil {
		t.Error("Esperava um erro para dados vazios, mas nenhum foi retornado.")
	}
}

func TestParseSelic(t *testing.T) {
	// Teste com dados válidos
	data := []map[string]interface{}{
		{"valor": "5.2"},
	}

	result, err := util.ParseSelic(data)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	expectedResult := 5.2
	if result != expectedResult {
		t.Errorf("Resultado inesperado. Esperado: %f, Obtido: %f", expectedResult, result)
	}

	// Teste com valor inválido
	_, err = util.ParseSelic([]map[string]interface{}{{"valor": "abc"}})
	if err == nil {
		t.Error("Esperava um erro para string inválida, mas nenhum foi retornado.")
	}

	// Teste com dados vazios
	_, err = util.ParseSelic([]map[string]interface{}{})
	if err == nil {
		t.Error("Esperava um erro para dados vazios, mas nenhum foi retornado.")
	}
}
