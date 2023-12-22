package display

import (
    "fmt"
    "profitability/cli/pkg/calculus"
)

func ShowPreResult(result calculus.ResultPre) {
    fmt.Printf("Pré Fixado \nLiq: %.2f%% \nIR: %.2f%%\n", result.ResultPre, result.TaxRate)
}

func ShowPosResult(result calculus.ResultPos) {
    fmt.Printf("Pós Fixado \nLiq: %.2f%% \nIR: %.2f%%\n", result.ResultPos, result.TaxRate)
}

func ShowIPCAResult(result calculus.ResultIPCA) {
    fmt.Printf("IPCA+ \nLiq: %.2f%% \nIR: %.2f%%\n", result.ResultIPCA, result.TaxRate)
}

func ShowPropResult(result calculus.ResultProp) {
    fmt.Printf("Proporcional: %.2f%% \nSelic Hoje: %.2f%%\n", result.ResultProp, result.Selic)
}
