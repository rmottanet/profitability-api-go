package integrations_test

import (
	"testing"

	"profitability/api/pkg/cache"
	"profitability/api/pkg/integrations"
)


func TestGetIPCAData(t *testing.T) {
	// Mock da cache
	mockCache := cache.NewRatesCache()

	// Chamada à função de integração
	ipcaData, err := integrations.GetIPCAData(mockCache)

	// Verificação de erro
	if err != nil {
		t.Errorf("Erro ao obter dados do IPCA: %v", err)
	}

	// Verificação da estrutura dos dados do IPCA
	if ipcaData == nil {
		t.Error("Dados do IPCA não devem ser nulos")
	}
	if ipcaData.Data == "" {
		t.Error("O campo de data do IPCA não deve estar vazio")
	}
	if ipcaData.Valor == 0 {
		t.Error("O valor do IPCA não deve ser zero")
	}

	// Verificação de armazenamento em cache
	cachedValue, exists := mockCache.Get("IPCA")
	if !exists {
		t.Error("IPCA deve ser armazenado em cache")
	}
	if cachedValue == 0 {
		t.Error("O valor armazenado em cache do IPCA não deve ser zero")
	}
}


