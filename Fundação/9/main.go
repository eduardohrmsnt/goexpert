package main

func main() {

	i := 0
	total := func() int {
		i += sum(1, 2, 3, 4, 5) * 2
		return i
	}

	// Exemplo de uso da função sum
	result := sum(1, 2, 3, 4, 5)

	println("A soma de 1, 2, 3, 4 e 5 é:", result) // Saída: A soma de 1, 2, 3, 4 e 5 é: 15
	println("O total é:", total)                   // Saída: O total é: 30
	println("A soma de 3 e 4 é:", result)          // Saída: A soma de 3 e 4 é: 7
	println(total())
	println(total())
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
