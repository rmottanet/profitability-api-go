package services_test

import (

	"testing"

	"profitability/api/pkg/services"
)

func TestPreService_Pre(t *testing.T) {
	// Criando uma instância do serviço de contratos pré-fixados
	preService := services.NewPreService()

	// Definindo valores de teste
	rate := 11.25
	term := 181

	// Executando o método Pre para calcular a rentabilidade líquida
	resultPre, taxRatePercentage, err := preService.Pre(rate, term)

	// Verificando se ocorreu algum erro durante o cálculo
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	// Definindo valores esperados
	expectedResultPre := 11.25 * (1 - 0.20)
	expectedTaxRatePercentage := 20.00

	// Verificando se os resultados estão corretos
	if resultPre != expectedResultPre {
		t.Errorf("Rentabilidade líquida incorreta. Esperado: %f, Obtido: %f", expectedResultPre, resultPre)
	}

	if taxRatePercentage != expectedTaxRatePercentage {
		t.Errorf("Taxa de imposto incorreta. Esperado: %f, Obtido: %f", expectedTaxRatePercentage, taxRatePercentage)
	}
}
