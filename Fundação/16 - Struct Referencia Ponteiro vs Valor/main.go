package main

import (
	"fmt"
	"runtime"
)

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	conta := &Conta{0}

	fmt.Printf("Criando conta com referência %p\n", conta)
	return conta
}

func NewContSemReferencia() Conta {
	conta := Conta{0}

	fmt.Printf("Criando conta sem referência %p\n", &conta)

	return conta
}

func (c *Conta) depositar(valor int) {
	c.saldo += valor
}

func (c *Conta) sacar(valor int) {
	c.saldo -= valor
}

func main() {
	conta2 := NewConta()
	fmt.Printf("Criando conta com referência %p\n", conta2)
	conta2.depositar(100)
	conta2.saldo = 200
	println(conta2.saldo)

	conta3 := NewContSemReferencia()
	fmt.Printf("Criando conta sem referência %p\n", &conta3)
	runtime.GC()
	conta3.depositar(100)

	fmt.Println("Saldo antes do depósito:", conta3.saldo)
	conta3.depositar(100)
	fmt.Println("Saldo depois do depósito:", conta3.saldo)
	conta3.sacar(50)
	fmt.Println("Saldo depois do saque:", conta3.saldo)
}
