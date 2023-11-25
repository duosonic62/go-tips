package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_DEFAULT, tablewriter.ALIGN_RIGHT})
	table.SetColumnColor(tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgBlueColor}, tablewriter.Colors{tablewriter.FgGreenColor})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
