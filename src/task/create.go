package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := Task{
			Id:          uuid.New(),
			Name:        Name,
			Description: Description,
			Owner:       Owner,
			Labels:      Labels,
			Status:      Status,
			CreatedAt:   time.Now(),
			Duration:    Duration,
		}

		tasks := LoadTasks()
		tasks = append(tasks, task)
		SaveTasks(tasks)

		fmt.Println("Task created successfully!")
	},
}
