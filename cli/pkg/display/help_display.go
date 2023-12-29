package display

import (
    "fmt"
)

func ShowUsage() {
	fmt.Println("Uso: profit <modalidade> <taxa> <prazo>")
	fmt.Println("Use 'profit --help' para obter mais informações.")
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

