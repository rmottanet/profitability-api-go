package util_test

import (
	"testing"

	"profitability/api/pkg/util"
)

func TestValidateRate(t *testing.T) {
	// Casos de teste para taxas válidas
	validRates := []string{"10", "10.2", "100.25", "0.5", "0.01"}
	for _, rate := range validRates {
		if !util.ValidateRate(rate) {
			t.Errorf("A validação falhou para a taxa %s, esperava-se que fosse válida", rate)
		}
	}

	// Casos de teste para taxas inválidas
	invalidRates := []string{"", "10..2", "abc", "-10", "0", "0.0", "1000.256"}
	for _, rate := range invalidRates {
		if util.ValidateRate(rate) {
			t.Errorf("A validação falhou para a taxa %s, esperava-se que fosse inválida", rate)
		}
	}
}

func TestValidateTerm(t *testing.T) {
	// Casos de teste para termos válidos
	validTerms := []string{"1", "10", "100", "9999"}
	for _, term := range validTerms {
		if !util.ValidateTerm(term) {
			t.Errorf("A validação falhou para o termo %s, esperava-se que fosse válido", term)
		}
	}

	// Casos de teste para termos inválidos
	invalidTerms := []string{"", "10.2", "abc", "-10", "0", "10000", "100000"}
	for _, term := range invalidTerms {
		if util.ValidateTerm(term) {
			t.Errorf("A validação falhou para o termo %s, esperava-se que fosse inválido", term)
		}
	}
}
