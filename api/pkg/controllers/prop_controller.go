package controllers

import (
    "strconv"
    "net/http"
    "encoding/json"

    "profitability/api/pkg/models"
    "profitability/api/pkg/services"
)


type PropController struct {
    Service services.PropServiceInterface
}


func NewPropController(service services.PropServiceInterface) *PropController {
    return &PropController{
        Service: service,
    }
}


// Lida com a requisição de cálculo proporcional à SELIC.
func (propCtrl *PropController) ProfitProp(w http.ResponseWriter, r *http.Request) {
	// Lida com os parâmetros da solicitação
    rateStr := r.URL.Query().Get("rate")
    
    rate, err := strconv.ParseFloat(rateStr, 64)
    if err != nil {
        http.Error(w, "Erro ao converter a taxa.", http.StatusBadRequest)
        return
    }

    // Obém o resultado usando o serviço    
    propResult, err := propCtrl.Service.Prop(rate)
    if err != nil {
        http.Error(w, "Erro ao calcular rendimento pré fixado.", http.StatusInternalServerError)
        return
    }

    // Cria a resposta com os valores obtidos do serviço
    response := models.NewProfitResponse("non tax", propResult, 0)

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error formatting response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}


