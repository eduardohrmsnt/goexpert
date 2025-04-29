package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	wesley := Cliente{
		Nome:  "Eduardo",
		Idade: 25,
		Ativo: true,
	}

	fmt.Printf("%+v\n", wesley)      // Saída: {Nome:Eduardo Idade:25 Ativo:true}
	fmt.Printf("%v\n", wesley)       // Saída: {Eduardo 25 true}
	fmt.Printf("%#v\n", wesley)      // Saída: main.Cliente{Nome:"Eduardo", Idade:25, Ativo:true}
	fmt.Printf("%T\n", wesley)       // Saída: main.Cliente
	fmt.Printf("%t\n", wesley.Ativo) // Saída: true
	fmt.Printf("%s\n", wesley.Nome)  // Saída: Eduardo
	fmt.Printf("%d\n", wesley.Idade) // Saída: 25

	wesley.Ativo = false
	fmt.Printf("%t\n", wesley.Ativo) // Saída: false
}
