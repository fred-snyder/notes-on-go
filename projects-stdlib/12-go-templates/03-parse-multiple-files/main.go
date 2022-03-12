package main

import (
	"log"
	"os"
	"text/template"
)

func e(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// parse a single file
	tpl, err := template.ParseFiles("template_01.gohtml") // tpl == 01
	tpl, err = template.ParseFiles("template_02.gohtml")  // tpl == 02 // overwrites tpl
	tpl, err = template.ParseFiles("template_01.gohtml")  // tpl == 01 // overwrites tpl
	e(err)

	// write to Stdout
	err = tpl.Execute(os.Stdout, nil) // tpl == 01
	e(err)

	// parse multiple files
	tpl, err = tpl.ParseFiles("template_02.gohtml", "template_03.gohtml", "template_04.gohtml")
	e(err)

	err = tpl.Execute(os.Stdout, nil) // Execute only executes the first parsed template // This will output "Template 01"
	e(err)

	// execute a specific template
	tpl.ExecuteTemplate(os.Stdout, "template_04.gohtml", nil)
}
