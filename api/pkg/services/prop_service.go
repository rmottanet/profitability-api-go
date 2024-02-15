package services

import (
	"log"
    "errors"

    "profitability/api/pkg/cache"
    "profitability/api/pkg/integrations"
)


type PropServiceInterface interface {
    Prop(rate float64) (float64, error)
}


type PropService struct {
    Cache cache.Cache
}


func NewPropService(cache cache.Cache) *PropService {
    return &PropService{
        Cache: cache,
    }
}


// Prop realiza o cálculo da taxa proporcioanl em relação a SELIC
func (propSrvc *PropService) Prop(rate float64) (float64, error) {
	// Recupera a taxa SELIC
    selicRate, ok := propSrvc.Cache.Get("SELIC")
    if !ok {
        if _, err := integrations.GetSELICData(propSrvc.Cache); err != nil {
            log.Printf("Erro ao obter dados da SELIC: %v", err)
            return 0, err
        }
        selicRate, ok = propSrvc.Cache.Get("SELIC")
        if !ok {
            return 0, errors.New("SELIC não encontrada no cache após a atualização")
        }
    }

    resultProp := rate * selicRate

    return resultProp, nil
}

