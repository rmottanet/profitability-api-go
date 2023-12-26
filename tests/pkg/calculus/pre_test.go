package calculus_test

import (
	"testing"

	"profitability/cli/pkg/calculus"
)

func TestPre(t *testing.T) {
	// Define os parâmetros de entrada para a função
	rate := 10.0
	term := 5

	// Chama a função que está sendo testada
	result, err := calculus.Pre(rate, term)

	// Verifica se ocorreu algum erro durante a execução da função
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	// Define o valor esperado com base nos parâmetros de entrada
	expectedResult := 10.0 * (1 - 0.225)
	
	// Compara o resultado retornado pela função com o resultado esperado
	if result.ResultPre != expectedResult {
		t.Errorf("Unexpected result. Expected: %.2f, got: %.2f", expectedResult, result.ResultPre)
	}
}
