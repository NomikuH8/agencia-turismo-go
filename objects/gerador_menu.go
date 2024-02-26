package objects

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type GeradorMenu struct{}

func (gm GeradorMenu) LimparTela() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (gm GeradorMenu) GerarMenu(caminho string, opcoes []string, temSair bool, sairText string) {
	if sairText == "" {
		sairText = "0) Retornar"
	}

	fmt.Println(caminho)
	fmt.Println()

	for i, opcao := range opcoes {
		fmt.Printf("%d) %s", i+1, opcao)
	}

	fmt.Println()

	if temSair {
		fmt.Println(sairText)
	}

	fmt.Println()
}

func (gm GeradorMenu) PedirInput(texto string) int {
	reader := bufio.NewReader(os.Stdin)

	input := ""
	numEscolhido := 0
	for input == "" {
		fmt.Println(texto)
		input, _ = reader.ReadString('\n')

		inputInteger, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Opção inválida.")
			input = ""
			continue
		}

		numEscolhido = inputInteger
	}

	return numEscolhido
}

func (gm GeradorMenu) PegarOpcaoEscolhida() {
}

func (gm GeradorMenu) FazerPerguntas() {
}
