package cep

type Cep struct {
	Cidade string `json:"localidade"`
	Rua    string `json:"logradouro"`
	Cep    string `json:"cep"`
}
