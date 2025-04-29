package main

import "fmt"


const a = "Hello World!"

type ID int

var (
	b bool
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	var meuArray [3]int

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 30

	fmt.Println(len(meuArray))
	fmt.Printf("O tipo de a é %v\n", meuArray)

	valor := meuArray[1:]

	fmt.Printf("O valor do meuArray[1:] é: %v\n", valor)

	for i := 0; i < len(meuArray); i++ {
		fmt.Printf("O valor do meuArray[%d] é: %d\n", i, meuArray[i])
		fmt.Println(meuArray[i])
	}
}
