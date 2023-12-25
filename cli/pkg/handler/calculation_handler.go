package handler

import (
	"profitability/cli/pkg/calculus"
	"profitability/cli/pkg/display"
)

func HandleProp(taxa float64) {
	result, err := calculus.Prop(taxa)
	ErrorHandler(err)
	display.ShowPropResult(result)
}

func HandlePre(taxa float64, prazo int) {
	result, err := calculus.Pre(taxa, prazo)
	ErrorHandler(err)
	display.ShowPreResult(result)
}

func HandlePos(taxa float64, prazo int) {
	result, err := calculus.Pos(taxa, prazo)
	ErrorHandler(err)
	display.ShowPosResult(result)
}

func HandleIPCA(taxa float64, prazo int) {
	result, err := calculus.IPCA(taxa, prazo)
	ErrorHandler(err)
	display.ShowIPCAResult(result)
}
