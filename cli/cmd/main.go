package main

import (
	"os"
	"profitability/cli/pkg/handler"
)

func main() {
	handler.HandleEntry()
	
	// Extrai os argumentos
	modalidade := os.Args[1]
	handler.HandleModal(modalidade)

	// Valida e converte a taxa para um valor numérico
	taxa := handler.HandleTaxa(os.Args[2])

	// Realiza cálculos com base na modalidade
	if modalidade == "prop" {
		handler.HandleProp(taxa)
		
	} else {
		
		handler.HandleFinal()
		
		// Extrai e valida o terceiro argumento da linha de comando (prazo)
		prazo := handler.HandlePrazo(os.Args[3])
		
		// Realiza cálculos com base na modalidade selecionada
		switch modalidade {
		case "pre":
			handler.HandlePre(taxa, prazo)

		case "pos":
			handler.HandlePos(taxa, prazo)

		case "ipca":
			handler.HandleIPCA(taxa, prazo)
		}
	}
}


