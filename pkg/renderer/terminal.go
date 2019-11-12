package renderer

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// Terminal renders results to console
type Terminal struct{}

func (r Terminal) call(data [][]interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"DATE", "TEAMS", "", "PLACE"})
	for _, line := range data {
		t.AppendRow(line)
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}
