package cep

import (
	"encoding/json"
	"io"
	"net/http"
)

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")

	res, err := http.Get(`https://viacep.com.br/ws/` + cepParam + `/json/`)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var cepConvertido Cep

	json.Unmarshal(body, &cepConvertido)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cepConvertido)
}
