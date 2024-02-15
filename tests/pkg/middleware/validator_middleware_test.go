package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"profitability/api/pkg/middleware"
)

func TestValidateInput_ValidRate(t *testing.T) {
	// Define uma função de manipulador fictícia para testar
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Success"))
	})

	// Cria um request HTTP falsificado com um rate válido
	req := httptest.NewRequest("GET", "/api/pre?rate=11.25&term=181", nil)

	// Cria um gravador de resposta falsificado para capturar a resposta
	recorder := httptest.NewRecorder()

	// Cria um manipulador HTTP usando a função ValidateInput
	handler := middleware.ValidateInput(mockHandler)

	// Executa a solicitação HTTP falsificada usando o manipulador
	handler.ServeHTTP(recorder, req)

	// Verifica se a resposta está correta
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler retornou status incorreto: obteve %v, queria %v",
			status, http.StatusOK)
	}

	expectedResponse := "Success"
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Handler retornou resposta incorreta: obteve %v, queria %v",
			recorder.Body.String(), expectedResponse)
	}
}

func TestValidateInput_InvalidRate(t *testing.T) {
	// Define uma função de manipulador fictícia para testar
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Cria um request HTTP falsificado com um rate inválido
	req := httptest.NewRequest("GET", "/api/prop?rate=invalid&term=181", nil)

	// Cria um gravador de resposta falsificado para capturar a resposta
	recorder := httptest.NewRecorder()

	// Cria um manipulador HTTP usando a função ValidateInput
	handler := middleware.ValidateInput(mockHandler)

	// Executa a solicitação HTTP falsificada usando o manipulador
	handler.ServeHTTP(recorder, req)

	// Verifica se a resposta está correta
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Handler retornou status incorreto: obteve %v, queria %v",
			status, http.StatusBadRequest)
	}

	expectedResponse := "Parâmetro 'rate' inválido.\n"
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Handler retornou resposta incorreta: obteve %v, queria %v",
			recorder.Body.String(), expectedResponse)
	}
}
