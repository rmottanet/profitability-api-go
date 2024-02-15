package models

import (
    "time"
    "strconv"
)


type RatesResponse struct {
    API               string            `json:"api"`
    APIDocumentation  string            `json:"api_documentation"`
    Timestamp         string            `json:"timestamp"`
    Index             map[string]string `json:"index"`
    License           string            `json:"license"`
    TermsOfUse        string            `json:"terms_of_use"`
}


func NewRatesResponse(rates map[string]float64, timestamp time.Time) *RatesResponse {
    // formata dados
    formattedSELIC := strconv.FormatFloat(rates["SELIC"]*100, 'f', 2, 64) + "%"

    formattedIPCA := strconv.FormatFloat(rates["IPCA"], 'f', 2, 64) + "%"

    return &RatesResponse{
        API:              "Profitability",
        APIDocumentation: "https://rmottanet.gitbook.io/profitability",
        Index: map[string]string{
            "SELIC": formattedSELIC,
            "IPCA":  formattedIPCA,
        },
        Timestamp:  timestamp.Format(time.RFC3339),
        License:    "https://raw.githubusercontent.com/rmottanet/profitability/main/LICENSE",
        TermsOfUse: "https://rmottanet.gitbook.io/profitability/profitability/profitability-api-terms-of-use",
    }
}
