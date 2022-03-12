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

func writeToStdout() {
	// parse a template file
	tpl, err := template.ParseFiles("template_main.gohtml")
	e(err)

	// write to Stdout
	// tpl is a pointer to an object of type template (from package template)
	// in other words a reference to a memory address that holds a template
	err = tpl.ExecuteTemplate(os.Stdout, "template_main.gohtml", nil) // execute a specific template
	err = tpl.Execute(os.Stdout, nil)                                 // execute all templates in the template object

	log.Println(string(colorGreen), "Execute writeStdout() succesfull", string(colorReset)) // log also writes the datetime
}

func main() {
	writeToStdout()
}
