package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func e(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("template_01.gohtml"))
}

func main() {
	// string
	dataString := "Hello world"
	err := tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", dataString)
	e(err)

	// slice
	dataSlice := []int{3, 5, 7, 8}
	err = tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", dataSlice)
	e(err)

	// array
	dataArray := [5]string{"Hello", "Hi", "Hallo", "Hola", "Bonjour"}
	err = tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", dataArray)
	e(err)

	// map
	dataMap := map[string]string{"NL": "Nederland", "US": "United States", "FR": "France"}
	err = tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", dataMap)
	e(err)

	// struct
	dataStruct := struct { // anonymous struct
		fieldName1 string
		fieldName2 int
	}{"Year", 2020}
	err = tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", dataStruct)

	// slice of struct
	err = sliceOfStruct(tpl)
	e(err)

	// struct of slice of struct
	err = structSliceStruct(tpl)
	e(err)
}
