package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

func main() {
	t := template.New("").
		Funcs(template.FuncMap{
			"FormatDateTime": func(format string, d time.Time) string {
				if d.IsZero() {
					return ""
				}
				return d.Format(format)
			}})
	tmpl := `{{ FormatDateTime "2006年01月02日" .}}`
	t = template.Must(t.Parse(tmpl))
	err := t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
