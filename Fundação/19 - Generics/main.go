package main

type Number interface {
	~int | ~float64
}

type MyNumber int

func Soma(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}
func SomaIntefaceVazia(m map[string]interface{}) interface{} {
	var soma int
	for _, v := range m {
		soma += v.(int)
	}
	return soma
}

func SomaGeneric[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	// Exemplo de uso
	var x interface{} = 10

	m := map[string]interface{}{"a": x, "b": x, "c": x}
	n := map[string]int{"a": 10, "b": 20, "c": 30}
	p := map[string]float64{"a": 10.5, "b": 20.5, "c": 30.5}
	q := map[string]MyNumber{"a": 10, "b": 20, "c": 30}

	somaGeneric := SomaGeneric(p)
	println("Soma com map[string]float64:", somaGeneric)

	somaGenericMyNumber := SomaGeneric(q)
	println("Soma com map[string]MyNumber:", somaGenericMyNumber)

	somaGenericInt := SomaGeneric(n)
	println("Soma com map[string]interface{}:", somaGenericInt)

	soma := Soma(n)
	somaInterfaceVazia := SomaIntefaceVazia(m)

	println("Soma com map[string]int:", somaInterfaceVazia.(int))
	println("Soma:", soma)

	compara1 := Compara(10, 10.00)
	compara2 := Compara(10.5, 10.5)
	compara3 := Compara(10, 20)

	println("Compara 1:", compara1)
	println("Compara 2:", compara2)
	println("Compara 3:", compara3)
}
