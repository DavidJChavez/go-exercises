package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles(`tpl_file.gohtml`)
	if err != nil {
		log.Fatalln(err)
	}

	newFile, err := os.Create(`index.html`)
	if err != nil {
		log.Println(`Error creating file`, err)
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(newFile)

	err = tpl.Execute(newFile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
