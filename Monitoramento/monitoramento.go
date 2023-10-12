package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorCount = 3
const delay = 5

func main() {

	showTitle()

	// loop infinito
	for {
		showMenu()

		action := getAction()

		// não usa break
		switch action {
		case 1:
			startMonitor()
		case 2:
			fmt.Println("Logs")
		case 0:
			fmt.Println("Encerrando...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func showTitle() {
	fmt.Println("------------------------------------------------")
	fmt.Println("- Monitoramento")
	fmt.Println("------------------------------------------------")
}

func showMenu() {
	fmt.Println()

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Encerrar")

	fmt.Println()
}

func getAction() int {
	action := -1

	fmt.Print("Opção: ")

	// & ponteiro -> igual no c
	fmt.Scan(&action)

	fmt.Println()

	return action
}

func startMonitor() {
	fmt.Println("Monitorando...")
	fmt.Println()

	sites := []string{
		"https://www.alura.com.br/",
		"https://www.caelum.com.br/",
		"https://random-status-code.herokuapp.com/"}

	for try := 0; try < monitorCount; try++ {
		// foreach
		for index, site := range sites {
			pingSite(index+1, site)
		}

		if isNotLast(try) {
			time.Sleep(delay * time.Second)
		}

		fmt.Println()
	}

}

func pingSite(sequence int, site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site", sequence, "->", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", sequence, "->", site, "esta com problemas. Status Code:", response.StatusCode)
	}
}

func isNotLast(try int) bool {
	return try+1 < monitorCount
}
