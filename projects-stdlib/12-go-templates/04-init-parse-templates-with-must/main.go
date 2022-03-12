package main

import (
	"html/template"
	"log"
	"os"
)

// define tpl on package scope
var tpl *template.Template

func e(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	// tpl = template.New("")
	// tpl, err := tpl.ParseGlob("templates/*")
	// fmt.Println(tpl)
	// e(err)

	// Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil.
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", nil)
	e(err)
}
