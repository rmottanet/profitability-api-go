package handler

import (
	"profitability/cli/pkg/calculus"
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/display"
)

func HandleProp(taxa float64) {
	result, err := calculus.Prop(taxa)
	util.ErrorHandler(err)
	display.ShowPropResult(result)
}

func HandlePre(taxa float64, prazo int) {
	result, err := calculus.Pre(taxa, prazo)
	util.ErrorHandler(err)
	display.ShowPreResult(result)
}

func HandlePos(taxa float64, prazo int) {
	result, err := calculus.Pos(taxa, prazo)
	util.ErrorHandler(err)
	display.ShowPosResult(result)
}

func HandleIPCA(taxa float64, prazo int) {
	result, err := calculus.IPCA(taxa, prazo)
	util.ErrorHandler(err)
	display.ShowIPCAResult(result)
}
