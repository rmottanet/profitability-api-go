package services

import (
	"log"
	
    "profitability/api/pkg/cache"
    "profitability/api/pkg/integrations"
)


type RatesServiceInterface interface {
    GetRatesCached() map[string]float64
}


type RatesService struct {
    Cache cache.Cache
}


func NewRatesService(cache cache.Cache) *RatesService {
    return &RatesService{
        Cache: cache,
    }
}


// GetRatesCached retorna os índices em cache.
func (ratesSrvc *RatesService) GetRatesCached() map[string]float64 {
	// recupera o cache
    indexRates, _, err := ratesSrvc.Cache.GetAll()
    if err != nil {
        log.Printf("Erro ao obter as taxas no cache: %v", err)
        return nil
    }

    // requisição para apis externas
    if len(indexRates) == 0 || err != nil {
        if _, err := integrations.GetSELICData(ratesSrvc.Cache); err != nil {
            log.Fatalf("Erro ao obter dados da SELIC: %v", err)
        }
        if _, err := integrations.GetIPCAData(ratesSrvc.Cache); err != nil {
            log.Fatalf("Erro ao obter dados do IPCA: %v", err)
        }
        indexRates, _, err = ratesSrvc.Cache.GetAll()
        if err != nil {
            log.Printf("Erro ao obter os índices do cache após a atualização: %v", err)
            return nil
        }
    }

    return indexRates
}

