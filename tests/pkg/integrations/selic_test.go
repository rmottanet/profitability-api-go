package integrations_test

import (
	"testing"

	"profitability/api/pkg/cache"
	"profitability/api/pkg/integrations"
)


func TestGetSELICData(t *testing.T) {
	// Mock da cache
	mockCache := cache.NewRatesCache()

	// Chamada à função de integração
	selicData, err := integrations.GetSELICData(mockCache)

	// Verificação de erro
	if err != nil {
		t.Errorf("Erro ao obter dados da SELIC: %v", err)
	}

	// Verificação da estrutura dos dados da SELIC
	if selicData == nil {
		t.Error("Dados da SELIC não devem ser nulos")
	}
	if selicData.Data == "" {
		t.Error("O campo de data da SELIC não deve estar vazio")
	}
	if selicData.Valor == 0 {
		t.Error("O valor da SELIC não deve ser zero")
	}

	// Verificação de armazenamento em cache
	cachedValue, exists := mockCache.Get("SELIC")
	if !exists {
		t.Error("SELIC deve ser armazenado em cache")
	}
	if cachedValue == 0 {
		t.Error("O valor armazenado em cache da SELIC não deve ser zero")
	}
}
