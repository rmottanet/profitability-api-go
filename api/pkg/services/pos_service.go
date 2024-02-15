package services

import (
	"log"
    "errors"

	"profitability/api/pkg/util"
    "profitability/api/pkg/cache"
    "profitability/api/pkg/integrations"
)


type PosServiceInterface interface {
    Pos(rate float64, term int) (float64, float64, error)
}


type PosService struct {
    Cache cache.Cache
}


func NewPosService(cache cache.Cache) *PosService {
    return &PosService{
        Cache: cache,
    }
}


// Pos realiza o cálculo da rentabilidade líquida de contratos pós fixados com base na SELIC | CDI
func (posSrvc *PosService) Pos(rate float64, term int) (float64, float64, error) {
	// Recupera a taxa SELIC
    selicRate, ok := posSrvc.Cache.Get("SELIC")
    if !ok {
        if _, err := integrations.GetSELICData(posSrvc.Cache); err != nil {
            log.Printf("Erro ao obter dados da SELIC: %v", err)
            return 0, 0, err
        }
        selicRate, ok = posSrvc.Cache.Get("SELIC")
        if !ok {
            return 0, 0, errors.New("SELIC não encontrada no cache após a atualização")
        }
    }

	// Calcula o imposto
    taxRate, err := util.Aliquot(term)
    if err != nil {
        return 0, 0, errors.New("Erro ao calcular a taxa: " + err.Error())
    }

	// Clacula liquidez
    resultPos := (rate * selicRate) * (1 - taxRate)
    taxRatePercentage := taxRate * 100

    return resultPos, taxRatePercentage, nil
}

