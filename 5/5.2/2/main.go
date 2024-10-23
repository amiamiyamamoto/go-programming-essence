package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
}

func main() {
	tmpl := `{{.Name}}`
	t := template.Must(template.New("").Parse(tmpl))

	user := User{Name: "John Doe"}
	err := t.Execute(os.Stdout, user)

	if err != nil {
		log.Fatal(err)
	}
}
