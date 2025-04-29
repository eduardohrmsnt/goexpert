package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed maybe after the last.")
	defer fmt.Println("This will be printed last.")
	fmt.Println("Hello, World!")
	fmt.Println("This is a test of the defer statement.")
	fmt.Println("This will be printed first.")
	fmt.Println("This will be printed second.")
	fmt.Println("This will be printed third.")

	// A ordem de execução do defer é LIFO (Last In First Out)
	// Ou seja, o último defer a ser chamado é o primeiro a ser executado
	// Isso significa que o último defer a ser chamado será o primeiro
	// Sendo assim "This will be printed last." será o primeiro
	// E "This will be printed maybe after the last." será o último
	// Isso é útil para garantir que certos recursos sejam liberados
	// ou que certas ações sejam executadas antes que o programa termine
	// Por exemplo, se você estiver abrindo um arquivo, pode usar defer
	// para garantir que o arquivo seja fechado quando o programa terminar
	// Isso é útil para evitar vazamentos de memória ou outros problemas
	// de gerenciamento de recursos
}
