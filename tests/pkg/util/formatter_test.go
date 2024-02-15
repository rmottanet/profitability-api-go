package util_test

import (
	"testing"
	"profitability/api/pkg/util"
)

func TestParseFloat(t *testing.T) {
	// Caso de teste positivo
	s := "3.14"
	expected := 3.14
	got, err := util.ParseFloat(s)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if got != expected {
		t.Errorf("ParseFloat(%s) = %f; esperado %f", s, got, expected)
	}

	// Caso de teste para entrada inválida
	s = "invalid"
	_, err = util.ParseFloat(s)
	if err == nil {
		t.Errorf("ParseFloat(%s) deveria retornar um erro", s)
	}
}

func TestParseIPCA(t *testing.T) {
	// Caso de teste positivo
	data := []map[string]interface{}{
		{"valor": "1.5"},
		{"valor": "2.5"},
	}
	expected := 4.0
	got, err := util.ParseIPCA(data)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if got != expected {
		t.Errorf("ParseIPCA(%v) = %f; esperado %f", data, got, expected)
	}

	// Caso de teste para nenhum dado retornado
	data = []map[string]interface{}{}
	_, err = util.ParseIPCA(data)
	if err == nil {
		t.Errorf("ParseIPCA(%v) deveria retornar um erro", data)
	}

	// Caso de teste para valor não sendo uma string
	data = []map[string]interface{}{
		{"valor": 1.5},
	}
	_, err = util.ParseIPCA(data)
	if err == nil {
		t.Errorf("ParseIPCA(%v) deveria retornar um erro", data)
	}
}

func TestParseSelic(t *testing.T) {
	// Caso de teste positivo
	data := []map[string]interface{}{
		{"valor": "10"},
	}
	expected := 0.1
	got, err := util.ParseSelic(data)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if got != expected {
		t.Errorf("ParseSelic(%v) = %f; esperado %f", data, got, expected)
	}

	// Caso de teste para nenhum dado retornado
	data = []map[string]interface{}{}
	_, err = util.ParseSelic(data)
	if err == nil {
		t.Errorf("ParseSelic(%v) deveria retornar um erro", data)
	}

	// Caso de teste para valor não sendo uma string
	data = []map[string]interface{}{
		{"valor": 10},
	}
	_, err = util.ParseSelic(data)
	if err == nil {
		t.Errorf("ParseSelic(%v) deveria retornar um erro", data)
	}
}
