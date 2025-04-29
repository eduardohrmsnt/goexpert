package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	println("slice")
	numeros := []int{1, 2, 3, 4, 5}
	for _, numero := range numeros {
		println(numero)
	}

	i := 0

	for i <= 10 {
		println(i)
		i++
	}

	for {
		println(i)
		i++
		if i > 10 {
			break
		}
	}

	println("Hello, World!")
}
