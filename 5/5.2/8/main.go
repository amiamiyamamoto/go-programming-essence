package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t := template.Must(template.New("").ParseGlob("template/*.tmpl"))
	err := t.ExecuteTemplate(os.Stdout, "index", "これは本文です。")
	if err != nil {
		log.Fatal(err)
	}
}
