package api_test

import (
	"testing"

	"profitability/cli/pkg/api"
)

func TestGetUrl(t *testing.T) {
	// Teste para um índice válido
	validIndex := "ipca"
	expectedValidURL := "https://api.bcb.gov.br/dados/serie/bcdata.sgs.433/dados/ultimos/12?formato=json"
	validURL := api.GetUrl(validIndex)
	if validURL != expectedValidURL {
		t.Errorf("Para o índice válido '%s', esperava-se '%s', mas obteve '%s'", validIndex, expectedValidURL, validURL)
	}

	// Teste para um índice inválido
	invalidIndex := "inexistente"
	expectedInvalidURL := ""
	invalidURL := api.GetUrl(invalidIndex)
	if invalidURL != expectedInvalidURL {
		t.Errorf("Para o índice inválido '%s', esperava-se uma string vazia, mas obteve '%s'", invalidIndex, invalidURL)
	}
}
