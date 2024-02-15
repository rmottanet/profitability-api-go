package services

import (
    "errors"

    "profitability/api/pkg/util"
)


type PreServiceInterface interface {
    Pre(rate float64, term int) (float64, float64, error)
}


type PreService struct {
}


func NewPreService() *PreService {
    return &PreService{}
}


// Pre realiza o cálculo da rentabilidade líquida de contratos pré fixados.
func (preSrvc *PreService) Pre(rate float64, term int) (float64, float64, error) {
	// Calcula o imposto
    taxRate, err := util.Aliquot(term)
    if err != nil {
        return 0, 0, errors.New("Erro ao calcular a taxa: " + err.Error())
    }

	// Clacula liquidez
    resultPre := rate * (1 - taxRate)
    taxRatePercentage := taxRate * 100

    return resultPre, taxRatePercentage, nil
}
