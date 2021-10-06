package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		nome   string  = "Matheus"
		versao float32 = 1.1
		opcao  int
	)
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa esta na versão", versao)

	menu()

	opcao = leOpcao()

	switch opcao {
	case 1:
		fmt.Println("Monitoramento")
	case 2:
		fmt.Println("Logs")
	case 0:
		fmt.Println("Esta deixando o Sistema")
		os.Exit(0)
	default:
		fmt.Println("Essa opção não existe")
	}
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func leOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

func monitoramento() {

}
