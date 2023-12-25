package handler

import (
	"os"
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/display"
)

func HandleEntry() {
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
}

func HandleModal(modalidade string) {
	if !util.IsValidModalidade(modalidade) {
		display.ShowModal()
		return
	}
}

func HandleTaxa(taxaStr string) (float64) {
	taxa, err := util.IsValidTaxa(taxaStr)
	util.ErrorHandler(err)
	return taxa
}

func HandlePrazo(prazoStr string) (int) {
	prazo, err := util.IsValidPrazo(prazoStr)
	util.ErrorHandler(err)
	return prazo
}

func HandleFinal() {
		// Para outras modalidades, verifica se o número correto de argumentos é fornecido
		if len(os.Args) != 4 {
			display.ShowUsage()
			return
		}
}
