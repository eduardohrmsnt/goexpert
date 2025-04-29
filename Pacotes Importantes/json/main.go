package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	NumeroConta int    `json:"n"`
	Titular     string `json:"t"`
}

func main() {
	conta := Conta{
		NumeroConta: 123456,
		Titular:     "João Silva",
	}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}
	println(string(res))
	// Retorno em JSON
	// {"NumeroConta":123456,"Titular":"João Silva"}

	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n":123456,"t":"Eduardo Silva"}`)

	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		panic(err)
	}
	
	println("Número da conta:", conta2.NumeroConta)
	println("Titular da conta:", conta2.Titular)
}
