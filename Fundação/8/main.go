package main

import "errors"

func main() {
	// Exemplo de uso da função sum
	result := sum(3, 4)
	println("A soma de 3 e 4 é:", result) // Saída: A soma de 3 e 4 é: 7

	// Exemplo de uso da função sum1
	result1 := sum1(5, 6)
	println("A soma de 5 e 6 é:", result1) // Saída: A soma de 5 e 6 é: 11

	// Exemplo de uso da função sum2
	result2, ok := sum2(-3, 8)
	if ok {
		println("A soma de 7 e 8 é:", result2) // Saída: A soma de 7 e 8 é: 15
	} else {
		println("Um ou mais números são negativos.")
	}

	// Exemplo de uso da função sum3
	result3, err := sum3(-4, 10)
	if err != nil {
		println(err.Error()) // Saída: um ou mais números são negativos
	} else {
		println("A soma de 9 e 10 é:", result3) // Saída: A soma de 9 e 10 é: 19
	}

	meuNome := meuNomeE("João")
	println(meuNome) // Saída: Meu nome é João

	meuNomeSobrenome := meuNomeESobrenomeSao("João", "Silva")

	println(meuNomeSobrenome) // Saída: Meu nome é João Silva
}

func sum(a int, b int) int {
	return a + b
}

func sum1(a, b int) int {
	return a + b
}

func sum2(a, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}

	return a + b, true
}

func sum3(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("um ou mais números são negativos")
	}

	return a + b, nil
}

func meuNomeE(nome string) string {
	return "Meu nome é " + nome
}

func meuNomeESobrenomeSao(nome, sobrenome string) string {
	return "Meu nome é " + nome + " " + sobrenome
}
