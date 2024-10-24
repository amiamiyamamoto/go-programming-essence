package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tmpl := `<div>{{.}}</div>`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, template.HTML(`<b>1</b>`))
	if err != nil {
		log.Fatal(err)
	}

	tmlpjs := `<script>{{.}}</script>`
	tjs := template.Must(template.New("").Parse(tmlpjs))
	err = tjs.Execute(os.Stdout, template.JS(`alert("Hello, World!")`))
	if err != nil {
		log.Fatal(err)
	}
}
