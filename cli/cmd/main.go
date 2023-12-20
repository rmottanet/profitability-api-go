package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"profitability/cli/pkg/calculus"
)

// Função para validar a modalidade
func isValidModalidade(modalidade string) bool {
	validModalidades := map[string]bool{
		"pre":  true,
		"pos":  true,
		"ipca": true,
		"prop": true,
	}

	return validModalidades[modalidade]
}

func main() {
	if len(os.Args) < 3 && (len(os.Args) != 2 && os.Args[1] != "--help") {
		fmt.Println("Usage: profit <modalidade> <taxa> [prazo] ")
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		showHelp()
		return
	}

	modalidade := os.Args[1]
	taxaStr := os.Args[2]

	// Verificar se a modalidade é válida
	if !isValidModalidade(modalidade) {
		log.Fatal("Modalidade Inválida. Escolha entre pre, pos, ipca ou prop.")
	}

	taxa, err := strconv.ParseFloat(taxaStr, 64)
	if err != nil {
		log.Fatal("Taxa Inválida:", err)
	}

	// Verificar se a modalidade é "prop" e chamar a rota correspondente
	if modalidade == "prop" {
		result, err := calculus.Prop(taxa)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	
		fmt.Println("Result:", result.ResultProp)
		fmt.Println("Selic:", result.Selic)
	} else {
		// Se a modalidade não for "prop", verificar o prazo e chamar a rota correspondente
		if len(os.Args) != 4 {
			log.Fatal("Usage: profit <modalidade> <taxa> <prazo>")
		}

		prazoStr := os.Args[3]
		prazo, err := strconv.ParseInt(prazoStr, 10, 64)
		if err != nil {
			log.Fatal("Prazo Inválido:", err)
		}

		if prazo <= 0 {
			log.Fatal("Prazo deve ser maior que zero")
		}

		fmt.Println("Modalidade:", modalidade)

		switch modalidade {
		case "pre":
			result, err := calculus.Pre(taxa, int(prazo))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result.ResultPre)
			fmt.Println("Tax Rate:", result.TaxRate)

		case "pos":
			result, err := calculus.Pos(taxa, int(prazo))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result.ResultPos)
			fmt.Println("Tax Rate:", result.TaxRate)

		case "ipca":
			result, err := calculus.IPCA(taxa, int(prazo))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result:", result.ResultIPCA)
			fmt.Println("Tax Rate:", result.TaxRate)
		}
	}
}

func showHelp() {
	fmt.Println("usage: profit <modalidade> <taxa> <prazo> ")
	fmt.Println("Parametros: ")
	fmt.Println("Modalidades: Digite, pre, pos, ipca ou prop para escolher a modalidade de cálculo ")
	fmt.Println("Taxa: Digite a taxa ao ano usando pontos como separador. Ex: 11.62 ")
	fmt.Println("Prazo: (Opcional) Digite o prazo em dias conforme anúnciado no contrato. Ex: 721 ")
	fmt.Println("  --help   Show this help message")
}
