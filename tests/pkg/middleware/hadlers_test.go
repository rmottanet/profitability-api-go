package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"profitability/api/pkg/middleware"
)

func TestProfitHandler(t *testing.T) {
	// Defina uma função de manipulador fictícia para testar
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ProfitHandlerFunc foi chamada com sucesso"))
	}

	// Crie um request HTTP falsificado
	req := httptest.NewRequest("GET", "/rates", nil)

	// Crie um gravador de resposta falsificado para capturar a resposta
	recorder := httptest.NewRecorder()

	// Crie um manipulador HTTP usando a função ProfitHandler
	handler := middleware.ProfitHandler(mockHandler)

	// Execute a solicitação HTTP falsificada usando o manipulador
	handler.ServeHTTP(recorder, req)

	// Verifique se a resposta está correta
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler retornou status incorreto: obteve %v, queria %v",
			status, http.StatusOK)
	}

	expectedResponse := "ProfitHandlerFunc foi chamada com sucesso"
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Handler retornou resposta incorreta: obteve %v, queria %v",
			recorder.Body.String(), expectedResponse)
	}
}
