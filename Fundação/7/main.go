package main

func main() {
	salarios := map[string]float64{
		"João":     1500.00,
		"Maria":    2000.00,
		"José":     2500.00,
		"Pedro":    3000.00,
		"Lucas":    3500.00,
		"Paula":    4000.00,
		"Fernanda": 4500.00,
		"Roberto":  5000.00,
	}

	for nome, salario := range salarios {
		if salario > 3000.00 {
			delete(salarios, nome)
		}
	}
	for nome, salario := range salarios {
		println(nome, salario)
	}
	// Saída:
	// João 1500
	// Maria 2000
	// José 2500
	// Paula 4000
	// Fernanda 4500
	// Roberto 5000
	// O loop não imprime os nomes e salários de João, Maria e José, pois eles foram removidos do mapa.
	// O loop imprime os nomes e salários de Paula, Fernanda e Roberto, pois eles não foram removidos do mapa.
}
