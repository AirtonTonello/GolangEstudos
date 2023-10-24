package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorCount = 3
const delay = 5
const logFilePath = "C:\\Users\\airto\\go\\src\\GolangEstudos\\Monitoramento\\log.txt"

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
			showLogs()
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

	sites := readFileSites()

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
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", sequence, "->", site, "foi carregado com sucesso!")
		logStatus(site, true)
	} else {
		fmt.Println("Site", sequence, "->", site, "esta com problemas. Status Code:", response.StatusCode)
		logStatus(site, false)
	}
}

func isNotLast(try int) bool {
	return try+1 < monitorCount
}

func readFileSites() []string {
	var sites []string

	// file, err := os.ReadFile("sites.txt")
	file, err := os.Open("C:\\Users\\airto\\go\\src\\GolangEstudos\\Monitoramento\\sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func logStatus(site string, status bool) {
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo de log:", err)
	}

	logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")

	logFile.Close()
}

func showLogs() {
	fmt.Println()
	fileLog, err := os.ReadFile(logFilePath)

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo de log:", err)
	}

	fmt.Println(string(fileLog))
}
