package display

import (
    "fmt"
    "profitability/cli/pkg/calculus"
)

func ShowPreResult(result calculus.ResultPre) {
    fmt.Printf("Pré Fixado: \nLiq: %.2f%% \nIR: %.2f%%\n", result.ResultPre, result.TaxRate)
}

func ShowPosResult(result calculus.ResultPos) {
    fmt.Printf("Pós Fixado: \nLiq: %.2f%% \nIR: %.2f%%\n", result.ResultPos, result.TaxRate)
}

func ShowIPCAResult(result calculus.ResultIPCA) {
    fmt.Printf("IPCA+ \nLiq: %.2f%% \nIR: %.2f%%\n", result.ResultIPCA, result.TaxRate)
}

func ShowPropResult(result calculus.ResultProp) {
    fmt.Printf("Proporcional: %.2f%% \nSelic Hoje: %.2f%%\n", result.ResultProp, result.Selic)
}

func ShowUsage() {
	fmt.Println("Uso: profit <modalidade> <taxa> <prazo>")
	fmt.Println("Use 'profit --help' para obter mais informações.")
}


func ShowModal() {
	fmt.Println("Modalidade Inválida. Escolha entre pre, pos, ipca ou prop.")
}

func ShowHelp() {
	fmt.Println("Uso: profit <modalidade> <taxa> <prazo>")
	fmt.Println("\nParametros: ")
	fmt.Println("   Modalidades: Digite, pre, pos, ipca ou prop para escolher a modalidade de cálculo ")
	fmt.Println("   Taxa: Digite a taxa ao ano usando pontos como separador. Ex: 11.62 ")
	fmt.Println("   Prazo: (Opcional) Digite o prazo em dias conforme anúnciado no contrato. Ex: 721")
	fmt.Println("\n   --help   Mosta esta mensagem.")
}

func ShowVersion() {
	fmt.Println("proft version 0.2")
}
