package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, cep := range os.Args[1:] {
		println(cep)
		res, err := http.Get(`https://viacep.com.br/ws/` + cep + `/json/`)
		if err != nil {
			fmt.Fprintf(os.Stderr, `Erro ao fazer requisição %v\n`, err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var cep Cep

		err = json.Unmarshal(body, &cep)

		if err != nil {
			fmt.Fprintf(os.Stderr, `Erro ao fazer requisição %v\n`, err)
		}

		fmt.Println(cep.Localidade)

		arquivo, err := os.Create(cep.Cep + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, `Erro ao criar arquivo %v\n`, err)
		}
		defer arquivo.Close()
		_, err = arquivo.WriteString(fmt.Sprintf(`{"cep":"%s","logradouro":"%s","complemento":"%s","unidade":"%s","bairro":"%s","localidade":"%s","uf":"%s","estado":"%s","regiao":"%s","ibge":"%s","gia":"%s","ddd":"%s","siafi":"%s"}`, cep.Cep, cep.Logradouro, cep.Complemento, cep.Unidade, cep.Bairro, cep.Localidade, cep.Uf, cep.Estado, cep.Regiao, cep.Ibge, cep.Gia, cep.Ddd, cep.Siafi))
		if err != nil {
			fmt.Fprintf(os.Stderr, `Erro ao escrever arquivo %v\n`, err)
		}
	}
}
