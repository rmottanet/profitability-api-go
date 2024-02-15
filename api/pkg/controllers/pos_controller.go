package controllers

import (
    "strconv"
    "net/http"
    "encoding/json"

    "profitability/api/pkg/models"
    "profitability/api/pkg/services"
)


type PosController struct {
    Service services.PosServiceInterface
}


func NewPosController(service services.PosServiceInterface) *PosController {
    return &PosController{
        Service: service,
    }
}


// Lida com a requisição de contratos pós fixados com referência à SELIC | CDI.
func (posCtrl *PosController) ProfitPos(w http.ResponseWriter, r *http.Request) {
	// Lida com os parâmetros da solicitação
    rateStr := r.URL.Query().Get("rate")
    termStr := r.URL.Query().Get("term")
    
    rate, err := strconv.ParseFloat(rateStr, 64)
    if err != nil {
        http.Error(w, "Erro ao converter a taxa.", http.StatusBadRequest)
        return
    }

    term, err := strconv.Atoi(termStr)
    if err != nil {
        http.Error(w, "Erro ao converter o prazo.", http.StatusBadRequest)
        return
    }
    
    // Obém o resultado usando o serviço
    posResult, taxRatePercentage, err := posCtrl.Service.Pos(rate, term)
    if err != nil {
        http.Error(w, "Erro ao calcular rendimento pré fixado.", http.StatusInternalServerError)
        return
    }

    // Cria a resposta com os valores obtidos do serviço
	response := models.NewProfitResponse("Floating Rate", posResult, taxRatePercentage)

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error formatting response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
