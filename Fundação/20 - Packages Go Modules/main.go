package main

import (
	"fmt"

	"github.com/eduardohrmsnt/curso-go/matematica"
	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)

	uuid := uuid.New()
	fmt.Println("UUID gerado:", uuid)
	fmt.Println("Valor soma", soma)
}
