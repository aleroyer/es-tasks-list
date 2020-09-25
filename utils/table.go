package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// PrintTasks will display your data into a table in CLI
func PrintTasks(data map[string]interface{}, options map[string]bool) {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	headers := table.Row{"node", "ip", "task_id", "duration"}
	if options["start_time"] {
		headers = append(headers, "start_time")
	}
	if options["cancellable"] {
		headers = append(headers, "cancellable")
	}
	if options["action"] {
		headers = append(headers, "action")
	}
	if options["description"] {
		headers = append(headers, "description")
	}
	t.AppendHeader(headers)

	for _, node := range data["nodes"].(map[string]interface{}) {
		nodeName := node.(map[string]interface{})["name"].(string)
		nodeIP := node.(map[string]interface{})["host"].(string)

		for taskID, task := range node.(map[string]interface{})["tasks"].(map[string]interface{}) {
			runningTime := task.(map[string]interface{})["running_time"].(string)
			row := table.Row{nodeName, nodeIP, taskID, runningTime}

			if options["start_time"] {
				startTime := task.(map[string]interface{})["start_time"].(string)
				row = append(row, startTime)
			}
			if options["cancellable"] {
				cancellable := task.(map[string]interface{})["cancellable"].(bool)
				row = append(row, cancellable)
			}
			if options["action"] {
				action := task.(map[string]interface{})["action"].(string)
				row = append(row, action)
			}
			if options["description"] {
				description := task.(map[string]interface{})["description"].(string)
				row = append(row, description)
			}

			t.AppendRow(row)
		}

	}

	t.Render()
}
