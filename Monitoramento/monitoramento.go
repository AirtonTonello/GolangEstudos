package main

import "fmt"

func main() {
	fmt.Println("------------------------------------------------")
	fmt.Println("- Monitoramento")
	fmt.Println("------------------------------------------------")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Encerrar")

	// & ponteiro -> igual no c
	action := 0
	fmt.Scan(&action)
}
