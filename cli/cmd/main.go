package main

import (
	"os"
	"profitability/cli/pkg/handler"
)

func main() {
	handler.HandleEntry()
	
	// Extrai os argumentos
	functionality := os.Args[1]
	handler.HandleModal(functionality)

	// Valida e converte a rate para um valor numérico
	rate := handler.HandleRate(os.Args[2])

	// Realiza cálculos com base na modalidade
	if functionality == "prop" {
		handler.HandleProp(rate)
		
	} else {
		
		handler.HandleFinal()
		
		// Extrai e valida o terceiro argumento da linha de comando (term)
		term := handler.HandleTerm(os.Args[3])
		
		// Realiza cálculos com base na modalidade selecionada
		switch functionality {
		case "pre":
			handler.HandlePre(rate, term)

		case "pos":
			handler.HandlePos(rate, term)

		case "ipca":
			handler.HandleIPCA(rate, term)
		}
	}
}


