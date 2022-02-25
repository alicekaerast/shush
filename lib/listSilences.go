package lib

import (
	"github.com/fatih/color"
	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/rodaine/table"
)

func ListSilences(c *client.Alertmanager) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Comment", "Created By")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	silencesResult, _ := c.Silence.GetSilences(nil)
	silences := silencesResult.GetPayload()
	for _, n := range silences {
		if *n.Status.State == models.SilenceStatusStateActive {
			tbl.AddRow(*n.ID, *n.Comment, *n.CreatedBy)
		}
	}
	tbl.Print()
}
