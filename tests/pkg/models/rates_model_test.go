// // profitability/tests/pkg/models/rates_model_test.go
package models_test

import (
	"testing"
	"time"

	"profitability/api/pkg/models" // Importe o pacote que contém o código que você deseja testar
)

func TestNewRatesResponse(t *testing.T) {
	// Dados de exemplo
	rates := map[string]float64{
		"SELIC": 0.045,
		"IPCA":  4.340001,
	}
	timestamp := time.Now()

	// Cria uma nova resposta de taxas
	response := models.NewRatesResponse(rates, timestamp)

	// Verifica os campos da resposta
	if response.API != "Profitability" {
		t.Errorf("API esperada: Profitability, obtida: %s", response.API)
	}

	if response.APIDocumentation != "https://rmottanet.gitbook.io/profitability" {
		t.Errorf("Documentação da API incorreta: %s", response.APIDocumentation)
	}

	// Verifica se o campo de tempo está no formato correto
	_, err := time.Parse(time.RFC3339, response.Timestamp)
	if err != nil {
		t.Errorf("Erro ao analisar o campo de tempo: %v", err)
	}

	if response.Index["SELIC"] != "4.50%" {
	    t.Errorf("SELIC incorreto: esperado 4.50%%, obtido %s", response.Index["SELIC"])
	}
	
	if response.Index["IPCA"] != "4.34%" {
	    t.Errorf("IPCA incorreto: esperado 3.20%%, obtido %s", response.Index["IPCA"])
	}


	if response.License != "https://raw.githubusercontent.com/rmottanet/profitability/main/LICENSE" {
		t.Errorf("URL da licença incorreta: %s", response.License)
	}

	if response.TermsOfUse != "https://rmottanet.gitbook.io/profitability/profitability/profitability-api-terms-of-use" {
		t.Errorf("Termos de uso incorretos: %s", response.TermsOfUse)
	}
}
