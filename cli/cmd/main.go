package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 && (len(os.Args) != 2 && os.Args[1] != "--help") {
		fmt.Println("Usage: profit <modalidade> <taxa> <prazo> ")
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		showHelp()
		return
	}

	modalidade := os.Args[1] // Corrigido para usar uma string
	taxa, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal("Taxa Inválida:", err)
	}

	prazo, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal("Prazo Inválido:", err)
	}

	result := taxa + prazo

	fmt.Println("Result:", result, modalidade)
}

func showHelp() {
	fmt.Println("Usage: profit <modalidade> <taxa> <prazo> ")
	fmt.Println("Parametros: ")
	fmt.Println("Modalidades: Digite, pre, pos, ipca ou prop para escolher a modalidade de cálculo ")
	fmt.Println("Taxa: Digite a taxa ao ano usando pontos como separador. Ex: 11.62 ")
	fmt.Println("Prazo: Digite o prazo em dias conforme anúnciado no contrato. Ex: 721 ")
	fmt.Println("  --help   Show this help message")
}
