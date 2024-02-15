package services

import (
	"log"
    "errors"

	"profitability/api/pkg/util"
    "profitability/api/pkg/cache"
    "profitability/api/pkg/integrations"
)


type IpcaServiceInterface interface {
    Ipca(rate float64, term int) (float64, float64, error)
}


type IpcaService struct {
    Cache cache.Cache
}


func NewIpcaService(cache cache.Cache) *IpcaService {
    return &IpcaService{
        Cache: cache,
    }
}


// Ipca realiza o cálculo da rentabilidade líquida de contratos pós fixados com base no IPCA.
func (ipcaSrvc *IpcaService) Ipca(rate float64, term int) (float64, float64, error) {
	// Recupera a taxa IPCA
    ipcaRate, ok := ipcaSrvc.Cache.Get("IPCA")
    if !ok {
        if _, err := integrations.GetIPCAData(ipcaSrvc.Cache); err != nil {
            log.Printf("Erro ao obter dados da IPCA: %v", err)
            return 0, 0, err
        }
        ipcaRate, ok = ipcaSrvc.Cache.Get("IPCA")
        if !ok {
            return 0, 0, errors.New("IPCA não encontrada no cache após a atualização")
        }
    }

	// Calcula o imposto
    taxRate, err := util.Aliquot(term)
    if err != nil {
        return 0, 0, errors.New("Erro ao calcular a taxa: " + err.Error())
    }

	// Clacula liquidez
    resultIpca := (rate + ipcaRate / 100) * (1 - taxRate)
    taxRatePercentage := taxRate * 100

    return resultIpca, taxRatePercentage, nil
}

