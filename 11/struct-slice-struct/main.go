package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type amigo struct {
	Nombre string
	Apodo  string
}

type carro struct {
	Modelo  string
	Puertas int
}

type item struct {
	Compas      []amigo
	Transportes []carro
}

func init() {
	tpl = template.Must(template.ParseFiles(`tpl.gohtml`))
}

func main() {
	rebe := amigo{
		Nombre: `Rebeca`,
		Apodo:  `La Rebechocha`,
	}
	xime := amigo{
		Nombre: `Ximena`,
		Apodo:  `Manis`,
	}
	dama := amigo{
		Nombre: `Damaris`,
		Apodo:  `Manis de la Rebechocha`,
	}

	vochido := carro{
		Modelo:  `VW`,
		Puertas: 2,
	}
	peugeot := carro{
		Modelo:  `Peugeot`,
		Puertas: 4,
	}

	compas := []amigo{xime, rebe, dama}
	chuchus := []carro{vochido, peugeot}

	data := item{
		Compas:      compas,
		Transportes: chuchus,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
