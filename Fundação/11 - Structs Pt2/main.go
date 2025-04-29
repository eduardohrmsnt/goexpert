package main

import "fmt"

type Endereco struct {
	Rua    string
	Cidade string
	Estado string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco
}

func main() {
	wesley := Cliente{
		Nome:  "Eduardo",
		Idade: 25,
		Ativo: true,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}
	fmt.Printf("%+v\n", wesley)                // Saída: {Nome:Eduardo Idade:25 Ativo:true Endereco:{Rua:Rua das Flores Cidade:São Paulo Estado:SP}}
	fmt.Printf("%v\n", wesley)                 // Saída: {Eduardo 25 true {Rua das Flores São Paulo SP}}
	fmt.Printf("%#v\n", wesley)                // Saída: main.Cliente{Nome:"Eduardo", Idade:25, Ativo:true, Endereco:main.Endereco{Rua:"Rua das Flores", Cidade:"São Paulo", Estado:"SP"}}
	fmt.Printf("%T\n", wesley)                 // Saída: main.Cliente
	fmt.Printf("%t\n", wesley.Ativo)           // Saída: true
	fmt.Printf("%s\n", wesley.Nome)            // Saída: Eduardo
	fmt.Printf("%d\n", wesley.Idade)           // Saída: 25
	fmt.Printf("%s\n", wesley.Rua)             // Saída: Rua das Flores
	fmt.Printf("%s\n", wesley.Cidade)          // Saída: São Paulo
	fmt.Printf("%s\n", wesley.Estado)          // Saída: SP
	fmt.Printf("%+v\n", wesley.Endereco)       // Saída: {Rua:Rua das Flores Cidade:São Paulo Estado:SP}
	fmt.Printf("%v\n", wesley.Endereco)        // Saída: {Rua das Flores São Paulo SP}
	fmt.Printf("%#v\n", wesley.Endereco)       // Saída: main.Endereco{Rua:"Rua das Flores", Cidade:"São Paulo", Estado:"SP"}
	fmt.Printf("%T\n", wesley.Endereco)        // Saída: main.Endereco
	fmt.Printf("%s\n", wesley.Endereco.Rua)    // Saída: Rua das Flores
	fmt.Printf("%s\n", wesley.Endereco.Cidade) // Saída: São Paulo
	fmt.Printf("%s\n", wesley.Endereco.Estado) // Saída: SP
}
