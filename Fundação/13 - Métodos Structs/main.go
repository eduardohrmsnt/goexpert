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

type Empresa struct {
	Nome string
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

func (e *Empresa) Desativar() {
	e.Nome = ""
}

type Pessoa interface {
	Desativar()
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

	desativarCliente(&wesley)

	empresa := Empresa{
		Nome: "Tech Solutions",
	}

	desativarCliente(&empresa)

	fmt.Printf("%+v\n", wesley) // Saída: {Nome:Eduardo Idade:25 Ativo:true Endereco:{Rua:Rua das Flores Cidade:São Paulo Estado:SP}}
}

func desativarCliente(p Pessoa) {
	p.Desativar()
}
