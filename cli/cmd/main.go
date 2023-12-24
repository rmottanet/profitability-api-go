package main

import (
	"os"
	"profitability/cli/pkg/calculus"
	"profitability/cli/pkg/util"
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

	// Valida a modalidade
	if !util.IsValidModalidade(modalidade) {
		display.ShowModal()
		return
	}

	// Valida e converte a taxa para um valor numérico
	taxaStr := os.Args[2]
	taxa, err := util.IsValidTaxa(taxaStr)
	util.ErrorHandler(err)

	// Realiza cálculos com base na modalidade
	if modalidade == "prop" {
		result, err := calculus.Prop(taxa)
		util.ErrorHandler(err)
		display.ShowPropResult(result)
		
	} else {
		
		// Para outras modalidades, verifica se o número correto de argumentos é fornecido
		if len(os.Args) != 4 {
			display.ShowUsage()
			return
		}
		
		// Extrai e valida o terceiro argumento da linha de comando (prazo)
		prazoStr := os.Args[3]
		prazo, err := util.IsValidPrazo(prazoStr)
		util.ErrorHandler(err)
		
		// Realiza cálculos com base na modalidade selecionada
		switch modalidade {
		case "pre":
			result, err := calculus.Pre(taxa, prazo)
			util.ErrorHandler(err)
			display.ShowPreResult(result)

		case "pos":
			result, err := calculus.Pos(taxa, prazo)
			util.ErrorHandler(err)
			display.ShowPosResult(result)

		case "ipca":
			result, err := calculus.IPCA(taxa, prazo)
			util.ErrorHandler(err)
			display.ShowIPCAResult(result)
		}
	}
}
