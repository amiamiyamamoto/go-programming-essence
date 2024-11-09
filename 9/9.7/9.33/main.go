package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very  very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "sign", "Rating"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
