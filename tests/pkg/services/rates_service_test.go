package services_test

import (
	"testing"

	"profitability/api/pkg/cache"
	"profitability/api/pkg/services"
)

func TestGetRatesCached(t *testing.T) {
	// Criando um cache simples para o teste
	mockCache := cache.NewRatesCache()

	// Criando um serviço de taxas com o cache mockado
	ratesService := services.NewRatesService(mockCache)

	// Chamando a função para obter as taxas em cache
	indexRates := ratesService.GetRatesCached()

	// Verificando se as taxas estão vazias no início
	if len(indexRates) != 0 {
		t.Errorf("Esperava-se que o cache estivesse vazio, mas foi encontrado %d itens", len(indexRates))
	}

	// Adicionando valores de mock ao cache
	mockCache.Set("SELIC", 5.0, 0)
	mockCache.Set("IPCA", 3.0, 0)

	// Chamando a função para obter as taxas em cache novamente
	indexRates = ratesService.GetRatesCached()

	// Verificando se as taxas foram adicionadas corretamente ao cache
	if len(indexRates) != 2 {
		t.Errorf("Esperava-se que o cache tivesse 2 itens, mas foram encontrados %d itens", len(indexRates))
	}

	// Verificando se as taxas foram corretamente armazenadas no cache
	if rate, ok := indexRates["SELIC"]; !ok || rate != 5.0 {
		t.Errorf("Valor incorreto para a taxa SELIC no cache")
	}

	if rate, ok := indexRates["IPCA"]; !ok || rate != 3.0 {
		t.Errorf("Valor incorreto para a taxa IPCA no cache")
	}
}
