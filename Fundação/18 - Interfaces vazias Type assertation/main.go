package main

func main() {
	var minhaVar interface{} = "Eduardo Hermes"

	println(minhaVar.(string))
	res, ok := minhaVar.(int)

	if ok {
		println(res)
	} else {
		println("Não é um inteiro")
	}
}
