package main

import "fmt"

type Endereco struct {
	Rua    string
	Cidade string
	Estado string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
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

	wesley.Desativar()
	fmt.Printf("%+v\n", wesley) // Saída: {Nome:Eduardo Idade:25 Ativo:true Endereco:{Rua:Rua das Flores Cidade:São Paulo Estado:SP}}
}
