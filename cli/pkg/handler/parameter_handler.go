package handler

import (
	"os"
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/display"
)

func HandleEntry() {
	if len(os.Args) < 2 {
		display.ShowUsage()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		display.ShowHelp()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "--version" {
		display.ShowVersion()
		os.Exit(0)
	}
}

func HandleModal(modalidade string) {
	if !util.IsValidModalidade(modalidade) {
		display.ShowModal()
		os.Exit(0)
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
	if len(os.Args) != 4 {
		display.ShowUsage()
		os.Exit(0)
	}
}
