package main

type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	switch t.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case float64:
		println("float64")
	default:
		println("unknown type")
	}
}
