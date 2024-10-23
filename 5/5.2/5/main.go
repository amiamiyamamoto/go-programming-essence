package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	type Employee struct {
		Name string
	}
	type Company struct {
		Employees []Employee
	}
	company := Company{
		Employees: []Employee{
			{Name: "John Doe"},
			{Name: "Mike"},
		},
	}

	tmpl := `{{with $v := index .Employees 0}}
	{{$v.Name}}
	{{else}}
	Not found
	{{end}}`

	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, company)
	if err != nil {
		log.Fatal(err)
	}
}
