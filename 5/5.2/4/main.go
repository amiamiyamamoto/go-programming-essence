package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	type User struct {
		Age  int
		Name string
	}
	tmpl := `{{if gt .Age 20}}
	{{.Name}} is older than 20
	{{else}}
	{{.Name}} is not older than 20
	{{end}}`

	user := User{Age: 21, Name: "John Doe"}
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, user)

	if err != nil {
		log.Fatal(err)
	}
}
