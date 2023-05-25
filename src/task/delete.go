package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := LoadTasks()

		if len(args) < 1 {
			fmt.Println("Please specify the task number to delete.")
			return
		}

		taskNum := Atoi(args[0])

		// if taskNum < 1 || taskNum > len(tasks) {
		// 	fmt.Println("Invalid task number.")
		// 	return
		// }

		tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

		SaveTasks(tasks)

		fmt.Println("Task deleted successfully!")
	},
}
