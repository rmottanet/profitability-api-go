package middleware

import "net/http"

type ProfitHandlerFunc func(http.ResponseWriter, *http.Request)


func ProfitHandler(fn ProfitHandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fn(w, r)
    })
}
