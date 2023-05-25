package task

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var FinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish a task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := LoadTasks()

		if len(args) < 1 {
			fmt.Println("Please specify the task number to finish.")
			return
		}

		for i, task := range tasks {
			if task.Id.String() == string(args[0]) {
				if task.Status == "finished" {
					fmt.Println("Task already finished.")
				} else {
					tasks[i].Status = "finished"
					tasks[i].FinishedAt = time.Now().Format("2006.01.02 15:04:05.000")
					SaveTasks(tasks)
					fmt.Println("Task finished successfully!")
				}
			}
		}
	},
}
