// profitability/tests/pkg/models/profit_model_test.go
package models_test

import (
	"testing"
	"time"

	"profitability/api/pkg/models" // Importe o pacote que contém o código que você deseja testar
)

func TestNewProfitResponse(t *testing.T) {
	// Parâmetros de exemplo
	tipo := "Example Type"
	profitResult := 5.75
	tribute := 1.5

	// Cria uma nova resposta de lucratividade
	response := models.NewProfitResponse(tipo, profitResult, tribute)

	// Verifica os campos da resposta
	if response.API != "Profitability" {
		t.Errorf("API esperada: Profitability, obtida: %s", response.API)
	}

	if response.APIDocumentation != "https://rmottanet.gitbook.io/profitability" {
		t.Errorf("Documentação da API incorreta: %s", response.APIDocumentation)
	}

	// Verifica se o campo de tempo está no formato correto
	_, err := time.Parse(time.RFC3339, response.TimeStamp)
	if err != nil {
		t.Errorf("Erro ao analisar o campo de tempo: %v", err)
	}

	if response.DebtInstruments != tipo {
		t.Errorf("Tipo de instrumentos de dívida incorreto: esperado %s, obtido %s", tipo, response.DebtInstruments)
	}

	if response.License != "https://raw.githubusercontent.com/rmottanet/profitability/main/LICENSE" {
		t.Errorf("URL da licença incorreta: %s", response.License)
	}

	if response.TermsOfUse != "https://rmottanet.gitbook.io/profitability/profitability/profitability-api-terms-of-use" {
		t.Errorf("Termos de uso incorretos: %s", response.TermsOfUse)
	}

	// Verifica se os campos tributos e lucros são exibidos corretamente quando os valores são diferentes de zero
	expectedLiquid := "5.75%"
	if response.Liquid != expectedLiquid {
	    t.Errorf("Lucro líquido incorreto: esperado %s, obtido %s", expectedLiquid, response.Liquid)
	}
	
	expectedTribute := "1.50%"
	if response.Tribute != expectedTribute {
	    t.Errorf("Tributo incorreto: esperado %s, obtido %s", expectedTribute, response.Tribute)
	}

}
