package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cep", BuscaCepHandler)
	mux.Handle("/cep-luiz-alves", &cepLuizAlves{"89128000"})
	http.ListenAndServe(":8080", mux)
}

type cepLuizAlves struct {
	Cep string
}

func (c *cepLuizAlves) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cepLA, err := BuscaCep(c.Cep)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cepLA)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cep" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request: Missing 'cep' parameter"))
		return
	}

	cep, err := BuscaCep(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cepParam string) (*Cep, error) {
	res, err := http.Get(`https://viacep.com.br/ws/` + cepParam + `/json/`)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var cep Cep
	err = json.Unmarshal(body, &cep)

	if err != nil {
		return nil, err
	}

	return &cep, nil
}
