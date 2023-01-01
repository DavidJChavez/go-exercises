package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type friend struct {
	Nombre string
	Apodo  string
}

func init() {
	tpl = template.Must(template.ParseFiles(`tpl.gohtml`))
}

func main() {
	rebe := friend{
		Nombre: `Rebeca`,
		Apodo:  `La Rebechocha`,
	}
	xime := friend{
		Nombre: `Ximena`,
		Apodo:  `Manis`,
	}
	dama := friend{
		Nombre: `Damaris`,
		Apodo:  `Manis de la Rebechocha`,
	}

	friends := []friend{rebe, xime, dama}

	err := tpl.Execute(os.Stdout, friends)
	if err != nil {
		log.Fatalln(err)
	}
}
