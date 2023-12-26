package util_test

import (
	"testing"
	"profitability/cli/pkg/util"
)

func TestIsValidModalidade(t *testing.T) {
	// Teste com modalidade válida
	result := util.IsValidModalidade("pre")
	if !result {
		t.Error("Esperava true para modalidade válida, mas obteve false.")
	}

	// Teste com modalidade inválida
	result = util.IsValidModalidade("invalida")
	if result {
		t.Error("Esperava false para modalidade inválida, mas obteve true.")
	}
}

func TestIsValidTaxa(t *testing.T) {
	// Teste com valor válido
	result, err := util.IsValidTaxa("10.5")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	expectedResult := 10.5
	if result != expectedResult {
		t.Errorf("Resultado inesperado. Esperado: %f, Obtido: %f", expectedResult, result)
	}

	// Teste com valor inválido
	_, err = util.IsValidTaxa("400")
	if err == nil {
		t.Error("Esperava um erro para taxa acima do limite, mas nenhum foi retornado.")
	}

	_, err = util.IsValidTaxa("abc")
	if err == nil {
		t.Error("Esperava um erro para string inválida, mas nenhum foi retornado.")
	}
}

func TestIsValidPrazo(t *testing.T) {
	// Teste com valor válido
	result, err := util.IsValidPrazo("365")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	expectedResult := 365
	if result != expectedResult {
		t.Errorf("Resultado inesperado. Esperado: %d, Obtido: %d", expectedResult, result)
	}

	// Teste com valor inválido
	_, err = util.IsValidPrazo("-5")
	if err == nil {
		t.Error("Esperava um erro para prazo negativo, mas nenhum foi retornado.")
	}

	_, err = util.IsValidPrazo("5000000000")
	if err == nil {
		t.Error("Esperava um erro para prazo muito grande, mas nenhum foi retornado.")
	}

	_, err = util.IsValidPrazo("abc")
	if err == nil {
		t.Error("Esperava um erro para string inválida, mas nenhum foi retornado.")
	}
}
