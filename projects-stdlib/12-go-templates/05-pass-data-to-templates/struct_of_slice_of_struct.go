package main

import (
	"html/template"
	"os"
)

type project struct {
	id          string // unique project ID
	housenumber int
}

type user struct {
	username string
	email    string
}

type bidding struct {
	project *project // takes a pointer to a project
	user    *user
	bidding int // amount in euro cents
}

// instead of explicitly defining a new templateData type
// you can also use an anonymous struct // see below
type templateData struct {
	projects []project
	users    []user
	biddings []bidding
}

func structSliceStruct(tpl *template.Template) error {

	proj1 := project{
		id:          "arp-1073ad-40",
		housenumber: 40,
	}

	user1 := user{
		username: "freek",
		email:    "freek@arpartners.nl",
	}

	bidding1 := bidding{
		project: &proj1, // the memory address of proj1
		user:    &user1, // the memory address of user1
		bidding: 300000,
	}

	projects := []project{proj1}
	users := []user{user1}
	// biddings := []bidding{bidding1}

	// instead of explicitly defining a new templateData type
	// you can also use an anonymous struct
	// data := struct{ projects []project }{projects: []project{proj1}}
	// data := struct{ projects []project }{[]project{proj1}}
	// data := struct{ projects []project }{proj1}
	data := templateData{
		projects: projects,
		users:    users,
		biddings: []bidding{bidding1}, // biddings
	}

	return tpl.ExecuteTemplate(os.Stdout, "template_01.gohtml", data)
}
