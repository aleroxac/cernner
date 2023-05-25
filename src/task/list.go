package task

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := LoadTasks()

		task_json_formated, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s\n", task_json_formated)
	},
}
