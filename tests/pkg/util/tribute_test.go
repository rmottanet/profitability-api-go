package util_test

import (
	"testing"

	"profitability/api/pkg/util"
)

func TestAliquot(t *testing.T) {
	// Casos de teste para prazos válidos
	testCases := []struct {
		termo    int
		expected float64
	}{
		{1, 0.225},   // 1 dia
		{90, 0.225},  // 90 dias
		{180, 0.225}, // 180 dias
		{181, 0.20},  // 181 dias
		{250, 0.20},  // 250 dias
		{360, 0.20},  // 360 dias
		{361, 0.175}, // 361 dias
		{540, 0.175}, // 540 dias
		{720, 0.175}, // 720 dias
		{721, 0.15}, // 721 dias
	}

	for _, tc := range testCases {
		got, err := util.Aliquot(tc.termo)
		if err != nil {
			t.Errorf("Para termo %d, esperava %f, mas obteve um erro: %v", tc.termo, tc.expected, err)
		}
		if got != tc.expected {
			t.Errorf("Para termo %d, esperava %f, mas obteve %f", tc.termo, tc.expected, got)
		}
	}

	// Casos de teste para prazos inválidos
	invalidTestCases := []int{0, -10}
	for _, termo := range invalidTestCases {
		_, err := util.Aliquot(termo)
		if err == nil {
			t.Errorf("Aliquot(%d) deveria retornar um erro, mas não retornou", termo)
		}
	}
}
