package main

import (
	cep "busca-cep/cep"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cep", cep.BuscaCep)

	http.ListenAndServe(":8080", mux)
}
