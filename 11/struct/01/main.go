package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type friend struct {
	Name  string
	Apodo string
}

func init() {
	tpl = template.Must(template.ParseFiles(`tpl.gohtml`))
}

func main() {
	rebe := friend{
		Name:  `Rebeca`,
		Apodo: `La Rebechocha`,
	}
	err := tpl.Execute(os.Stdout, rebe)
	if err != nil {
		log.Fatalln(err)
	}
}
