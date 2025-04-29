package main

func main() {
	//Memória => Endereço => Valor

	a := 10    //Abre uma área de memória e coloca o valor 10
	b := 20    //Abre outra área de memória e coloca o valor 20
	c := a + b //Abre outra área de memória e coloca o valor 30
	//a, b e c são variáveis que armazenam valores inteiros

	//É como se abrisse uma caixa do correio e colocasse o valor 10 dentro dela
	//Quando você quer acessar o valor, você vai até a caixa do correio e pega o valor que está dentro dela
	//Se você quiser mudar o valor, você vai até a caixa do correio e troca o valor que está dentro dela
	//Se você quiser saber o endereço da caixa do correio, você pode usar o operador &
	//O operador & retorna o endereço da caixa do correio
	//O operador * retorna o valor que está dentro da caixa do correio

	println(a) // Saída: 10
	println(b) // Saída: 20
	println(c) // Saída: 30

	//O operador & retorna o endereço da caixa do correio
	println(&a) // Saída: 0xc00000a0b8
	println(&b) // Saída: 0xc00000a0c0
	println(&c) // Saída: 0xc00000a0c8

	//O operador * retorna o valor que está dentro da caixa do correio
	println(*&a) // Saída: 10
	println(*&b) // Saída: 20
	println(*&c) // Saída: 30
	//O operador * também pode ser usado para desreferenciar um ponteiro
	//Desreferenciar um ponteiro significa acessar o valor que está dentro do ponteiro
	//O operador * é usado para desreferenciar um ponteiro
	//Para modificar o valor que está dentro do ponteiro, você pode usar o operador * e atribuir um novo valor

	var ponteiro *int = &a //Cria um ponteiro que aponta para a variável a
	//O operador & retorna o endereço da variável a
	//O ponteiro é uma variável que armazena o endereço da variável a

	ponteiro = &b //O ponteiro agora aponta para a variável b

	g := *ponteiro
	//O operador * retorna o valor que está dentro do ponteiro
	//g não é um ponteiro, é uma variável que armazena o valor que está dentro do ponteiro
	//g é uma cópia do valor que está dentro do ponteiro, porém é criado um novo espaço na memória

	e := &g

	*e = 50

	println(*e)
	println(g)
	println(*ponteiro)
	

	//A variavel aponta para o endereço da variável a que pode ser acessado com o operador *

	*ponteiro = 100 //Desreferencia o ponteiro e atribui o valor 100 à variável a

	println(a)
	println(&a)
	println(*ponteiro)
	println(ponteiro)

	//Referencing: & => Endereço
	//Dereferencing: * => Valor
	//O operador * é usado para desreferenciar um ponteiro
}
