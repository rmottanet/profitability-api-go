package main

import (
	"fmt"
	"log"
	"os"
	"profitability/cli/pkg/calculus"
	"profitability/cli/pkg/util"
	"profitability/cli/pkg/display"
)

func main() {
	if len(os.Args) < 3 && (len(os.Args) != 2 && os.Args[1] != "--help") {
		fmt.Println("Uso: profit <modalidade> <taxa> <prazo> ")
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		showHelp()
		return
	}

	modalidade := os.Args[1]
	taxaStr := os.Args[2]

	if !util.IsValidModalidade(modalidade) {
		log.Fatal("Modalidade Inválida. Escolha entre pre, pos, ipca ou prop.")
	}

	taxa, err := util.IsValidTaxa(taxaStr)
	if err != nil {
		log.Fatal("Erro: ", err)
	}

	// Verificar se a modalidade é "prop" e chamar a rota correspondente
	if modalidade == "prop" {
		result, err := calculus.Prop(taxa)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		
		display.ShowPropResult(result)
		
	} else {

		if len(os.Args) != 4 {
			log.Fatal("Uso: profit <modalidade> <taxa> <prazo>")
		}

		prazoStr := os.Args[3]
		
		prazo, err := util.IsValidPrazo(prazoStr)
		if err != nil {
			log.Fatal("Erro: ", err)
		}
	
		switch modalidade {
		case "pre":
			result, err := calculus.Pre(taxa, prazo)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			
			display.ShowPreResult(result)

		case "pos":
			result, err := calculus.Pos(taxa, prazo)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			
			display.ShowPosResult(result)

		case "ipca":
			result, err := calculus.IPCA(taxa, prazo)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			
			display.ShowIPCAResult(result)
		}
	}
}

func showHelp() {
	fmt.Println("Uso: profit <modalidade> <taxa> <prazo> \n")
	fmt.Println("Parametros: ")
	fmt.Println("   Modalidades: Digite, pre, pos, ipca ou prop para escolher a modalidade de cálculo ")
	fmt.Println("   Taxa: Digite a taxa ao ano usando pontos como separador. Ex: 11.62 ")
	fmt.Println("   Prazo: (Opcional) Digite o prazo em dias conforme anúnciado no contrato. Ex: 721 \n")
	fmt.Println("   --help   Show this help message")
}
