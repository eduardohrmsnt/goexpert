package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

// A função sum recebe dois ponteiros para inteiros e retorna a soma dos valores apontados por eles, alterando o valor do primeiro ponteiro para 50

func sum2(a, b int) int {
	b = 100
	return a + b
}

// A função sum2 recebe dois inteiros e retorna a soma deles porém não altera os valores originais, ela recebe uma cópia dos valores

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println(sum2(minhaVar1, minhaVar2))
	// A função sum2 não altera o valor de minhaVar1 e retorna a soma de 10 + 20
	println(sum(&minhaVar1, &minhaVar2))
	// A função sum altera o valor de minhaVar1 para 50 e retorna a soma de 50 + 20
	println(minhaVar1) // Saída: 10
}
