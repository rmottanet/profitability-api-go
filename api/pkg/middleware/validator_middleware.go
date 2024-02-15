package middleware

import (
    "log"
    "net/http"   
    "strings"
    
    "profitability/api/pkg/util"
)

func ValidateInput(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        routeName := strings.TrimPrefix(path, "/api/")

        // Validação do parâmetro 'rate'
        rate := r.URL.Query().Get("rate")
        if !util.ValidateRate(rate) {
            http.Error(w, "Parâmetro 'rate' inválido.", http.StatusBadRequest)
            return
        }

        // Validação do parâmetro 'term', se aplicável
        if routeName != "prop" {
            term := r.URL.Query().Get("term")
            log.Printf("Parâmetro 'term' recebido: %s", term)
            if !util.ValidateTerm(term) {
                http.Error(w, "Parâmetro 'term' inválido.", http.StatusBadRequest)
                return
            }
        }

        // Chama o próximo handler se todos os parâmetros estiverem corretos
        next.ServeHTTP(w, r)
    })
}

