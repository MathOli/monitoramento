package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	monitoramentos = 5
	delay          = 3
)

func main() {

	for {
		var (
			nome   string  = "Matheus"
			versao float32 = 1.3
			opcao  int
		)
		fmt.Println("Olá, sr.", nome)
		fmt.Println("Esse programa esta na versão", versao)

		menu()

		opcao = leOpcao()

		switch opcao {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Logs")
		case 0:
			fmt.Println("Esta deixando o Sistema")
			os.Exit(0)
		default:
			fmt.Println("Essa opção não existe")
			os.Exit(-1)
		}
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
	fmt.Println("")

	return opcao
}

func monitoramento() {
	fmt.Println("Monitoramento...")
	sites := []string{"https://www.youtube.com", "https://random-status-code.herokuapp.com"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Foi carregado com sucesso")
	} else {
		fmt.Println("Possue algum error: ", resp.StatusCode)
	}
}
