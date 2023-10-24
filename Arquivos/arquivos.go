package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	readFilePath := "C:\\Users\\airto\\go\\src\\GolangEstudos\\Arquivos\\data.txt"
	readFile(readFilePath)

	lines := []string{"linha 1", "linha 2", "linha 3", "linha 4", "linha 5"}
	writeFilePath := "C:\\Users\\airto\\go\\src\\GolangEstudos\\Arquivos\\data2.txt"
	writeLines(writeFilePath, lines)
}

func readFile(filePath string) {
	fmt.Println("Read File:", filePath)

	fileStream, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
	}

	reader := bufio.NewReader(fileStream)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		fmt.Println(line)

		if err == io.EOF {
			break
		}
	}

	fileStream.Close()
}

func writeLines(filePath string, lines []string) {
	fmt.Println("Write File:", filePath)

	// 0666 parece permissao do linux
	// concatenador de flag | -> igual no c#
	fileStream, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo data2:", err)
	}

	// writer := bufio.NewWriter(fileStream)

	for _, line := range lines {
		fmt.Println(line)
		line = line + "\n"
		fileStream.WriteString(line)
		// writer.WriteString(line)
	}

	// despesa buffer no arquivo
	// writer.Flush()

	fileStream.Close()
}
