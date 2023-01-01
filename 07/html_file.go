package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "David Rodriguez"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello Mom!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	newFile, err := os.Create(`index.html`)
	if err != nil {
		log.Fatal(`Error creating file`, err)
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(newFile)

	_, err = io.Copy(newFile, strings.NewReader(str))
	if err != nil {
		return
	}
}
