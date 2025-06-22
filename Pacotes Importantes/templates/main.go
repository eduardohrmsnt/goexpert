package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	tmp := template.New("template.html")

	tmp.Funcs(template.FuncMap{"toUpper": strings.ToUpper})

	tmp.ParseFiles("template.html")

	err := tmp.Execute(os.Stdout, Cursos{
		{"Aprendendo Go", 40},
		{"Aprendendo Python", 30},
		{"Aprendendo Java", 50},
	})

	if err != nil {
		panic(err)
	}
}
