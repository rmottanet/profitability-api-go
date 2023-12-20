package cmd_test

import (
	"os"
	"testing"

	"profitability/cli/cmd"
	"profitability/cli/pkg/calculus"
)

func TestMain(t *testing.T) {
	// Teste de chamada com o argumento --help
	os.Args = []string{"profit", "--help"}
	cmd.Main()

	// Teste de chamada com menos de três argumentos (exceto --help)
	os.Args = []string{"profit", "pre"}
	cmd.Main()

	// Teste de chamada com modalidade inválida
	os.Args = []string{"profit", "invalid", "10"}
	cmd.Main()

	// Teste de chamada com taxa inválida
	os.Args = []string{"profit", "pre", "invalid"}
	cmd.Main()

	// Teste de chamada com modalidade "prop"
	os.Args = []string{"profit", "prop", "10"}
	cmd.Main()

	// Teste de chamada com modalidade válida e prazo inválido
	os.Args = []string{"profit", "pre", "10"}
	main.Main()

	// Teste de chamada com modalidade válida e prazo válido
	os.Args = []string{"profit", "pre", "10", "5"}
	main.Main()

	// Teste de chamada com modalidade inválida e prazo inválido
	os.Args = []string{"profit", "invalid", "10", "5"}
	main.Main()
}
