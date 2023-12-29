package handler

import (
	"profitability/cli/pkg/calculus"
	"profitability/cli/pkg/display"
)

func HandleProp(rate float64) {
	result, err := calculus.Prop(rate)
	ErrorHandler(err)
	display.ShowPropResult(result)
}

func HandlePre(rate float64, term int) {
	result, err := calculus.Pre(rate, term)
	ErrorHandler(err)
	display.ShowPreResult(result)
}

func HandlePos(rate float64, term int) {
	result, err := calculus.Pos(rate, term)
	ErrorHandler(err)
	display.ShowPosResult(result)
}

func HandleIPCA(rate float64, term int) {
	result, err := calculus.IPCA(rate, term)
	ErrorHandler(err)
	display.ShowIPCAResult(result)
}
