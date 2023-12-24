package main

import (
	"os"
	"profitability/cli/pkg/handler"
	"profitability/cli/pkg/display"
)

func main() {
	// Verifica se há argumentos de linha de comando
	if len(os.Args) < 2 {
		display.ShowUsage()
		return
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		display.ShowHelp()
		return
	}

	if len(os.Args) == 2 && os.Args[1] == "--version" {
		display.ShowVersion()
		return
	}
	
	// Extrai os argumentos
	modalidade := os.Args[1]
	handler.HandleModal(modalidade)

	// Valida e converte a taxa para um valor numérico
	taxaStr := os.Args[2]
	taxa := handler.HandleTaxa(taxaStr)

	// Realiza cálculos com base na modalidade
	if modalidade == "prop" {
		handler.HandleProp(taxa)
		
	} else {
		
		// Para outras modalidades, verifica se o número correto de argumentos é fornecido
		if len(os.Args) != 4 {
			display.ShowUsage()
			return
		}
		
		// Extrai e valida o terceiro argumento da linha de comando (prazo)
		prazoStr := os.Args[3]
		prazo := handler.HandlePrazo(prazoStr)
		
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


