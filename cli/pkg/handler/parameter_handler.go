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

func HandleModal(functionality string) {
	if !util.IsValidFunctionality(functionality) {
		display.ShowModal()
		os.Exit(0)
	}
}

func HandleRate(rateStr string) (float64) {
	rate, err := util.IsValidRate(rateStr)
	ErrorHandler(err)
	return rate
}

func HandleTerm(termStr string) (int) {
	term, err := util.IsValidTerm(termStr)
	ErrorHandler(err)
	return term
}

func HandleFinal() {
	if len(os.Args) != 4 {
		display.ShowUsage()
		os.Exit(0)
	}
}
