package main

import (
	"html/template"
	"os"
)

type house struct {
	street      string
	housenumber int
	price       int
}

func sliceOfStruct(tpl *template.Template) error {

	h1 := house{
		street:      "Cuyp",
		housenumber: 22,
		price:       300000,
	}

	h2 := house{
		street:      "Cuyp",
		housenumber: 23,
		price:       400000,
	}

	h3 := house{
		street:      "Cuyp",
		housenumber: 25,
		price:       350000,
	}

	data := []house{h1, h2, h3}

	return tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", data)
}
