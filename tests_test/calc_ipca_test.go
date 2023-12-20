package calculus_test

import (
	"testing"
	"math"

	"profitability/cli/pkg/calculus"
)

func TestIPCA(t *testing.T) {
	// Define os parâmetros de entrada para a função
	rate := 10.0
	term := 5

	// Chama a função que está sendo testada
	result, err := calculus.IPCA(rate, term)

	// Verifica se ocorreu algum erro durante a execução da função
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Define o valor esperado com base nos parâmetros de entrada
	ipca := 4.59 // Mock do valor de IPCA
	taxRate := 0.225 // Mock da taxa de imposto

	expectedResult := (rate + ipca / 100) * (1 - taxRate)
	expectedResult = math.Round(expectedResult*100) / 100
	
	// Compara o resultado retornado pela função com o resultado esperado
	if result.ResultIPCA != expectedResult {
		t.Errorf("Unexpected result. Expected: %.2f, got: %.2f", expectedResult, result.ResultIPCA)
	}
}
