package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Olá Mundo dois tres quatro cinco seis sete oito nove dez"))

	if err != nil {
		panic(err)
	}
	println("Tamanho do arquivo: ", tamanho)

	defer f.Close()

	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 4)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	defer arquivo2.Close()

	os.Remove("arquivo.txt")
}
