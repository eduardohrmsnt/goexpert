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
	fmt.Printf("O tipo de a é %T\n", a)
	println(d)
	println(e)
	println(b)
	println(c)
	println(a)
	println(f)
}
