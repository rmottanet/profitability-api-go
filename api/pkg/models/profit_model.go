package models

import (
    "time"
    "strconv"
)


type ProfitRequest struct {
    Rate float64 `json:"rate"`
    Term int     `json:"term,omitempty"` 
}


type ProfitResponse struct {
    API               string  `json:"api"`
    APIDocumentation  string  `json:"api_documentation"`
    TimeStamp         string  `json:"time_stamp"`
    DebtInstruments   string  `json:"debt instruments"` 
    Liquid            string  `json:"liquid"`
    Tribute           string  `json:"tribute,omitempty"` 
    License           string  `json:"license"`
    TermsOfUse        string  `json:"terms_of_use"`
}


func NewProfitResponse(tipo string, profitResult, tribute float64) *ProfitResponse {
    response := &ProfitResponse{
        API:              "Profitability",
        APIDocumentation: "https://rmottanet.gitbook.io/profitability",
        TimeStamp:        time.Now().Format(time.RFC3339),
        DebtInstruments:  tipo,
        License:          "https://raw.githubusercontent.com/rmottanet/profitability/main/LICENSE",
        TermsOfUse:       "https://rmottanet.gitbook.io/profitability/profitability/profitability-api-terms-of-use",
        Liquid:           strconv.FormatFloat(profitResult, 'f', 2, 64) + "%",
    }

    if tribute > 0 {
        response.Tribute = strconv.FormatFloat(tribute, 'f', 2, 64) + "%"
    }

    return response
}

