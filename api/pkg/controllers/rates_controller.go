package controllers

import (
    "time"
    "net/http"
    "encoding/json"

    "profitability/api/pkg/models"
    "profitability/api/pkg/services"
)


type RatesController struct {
    Service services.RatesServiceInterface
}


func NewRatesController(service services.RatesServiceInterface) *RatesController {
    return &RatesController{
        Service: service,
    }
}


// Lida com a requisição de taxas disponíveis na api.
func (ratesCtrl *RatesController) GetRates(w http.ResponseWriter, r *http.Request) {
    // Obém o resultado usando o serviço
    rates := ratesCtrl.Service.GetRatesCached()

    // Cria a resposta com os valores obtidos do serviço
    response := models.NewRatesResponse(rates, time.Now())

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error formatting response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")    
    w.WriteHeader(http.StatusOK)    
    w.Write(jsonResponse)
}
