package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// tmpl := `{{range .}}
	// <p>{{.}}</p>{{end}}`
	tmpl := `{{index . 1}}`

	t := template.Must(template.New("").Parse(tmpl))
	values := []string{"Hello", "World"}
	err := t.Execute(os.Stdout, values)
	if err != nil {
		log.Fatal(err)
	}
}
