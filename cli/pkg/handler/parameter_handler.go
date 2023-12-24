package handler

import (
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/display"
)

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
