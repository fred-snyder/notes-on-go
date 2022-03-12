package main

import (
	"log"
	"os"
	"text/template"
)

const colorGreen string = "\033[32m"
const colorReset string = "\033[0m"

func e(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func writeToFile() {
	// parse a template file
	tpl, err := template.ParseFiles("template_main.gohtml")
	e(err)

	// create to a file
	f, err := os.Create("index.html")
	e(err)

	// DEFER Close()
	defer f.Close() // don't forget to Close() the file

	// execute and write to a file
	err = tpl.Execute(f, nil)
	e(err)

	log.Println(string(colorGreen), "Execute writeToFile() succesfull", string(colorReset))
}

func main() {
	writeToFile()
}
