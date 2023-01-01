package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(`templates/*.gomd`))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, `tpl_three.gomd`, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, `tpl_two.gomd`, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, `tpl_one.gomd`, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
