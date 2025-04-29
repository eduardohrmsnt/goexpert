package main

func main() {

	// Condicionais
	// if, else if, else
	if true {
		println("Verdadeiro")
	} else {
		println("Falso")
	}

	a := 10
	b := 20

	if a > b {
		println("A é maior que B")
	}
	if a < b {
		println("A é menor que B")
	}
	if a == b {
		println("A é igual a B")
	}
	if a != b {
		println("A é diferente de B")
	}
	if a >= b {
		println("A é maior ou igual a B")
	}
	if a <= b {
		println("A é menor ou igual a B")
	}
	if a == 10 && b == 20 {
		println("A é 10 e B é 20")
	}
	if a == 10 || b == 30 {
		println("A é 10 ou B é 30")
	}
	if !(a == 10) {
		println("A não é 10")
	}

	// Switch
	switch a {
	case 10:
		println("A é 10")
	case 20:
		println("A é 20")
	case 30:
		println("A é 30")
	default:
		println("A não é 10, 20 ou 30")
	}
}
