package task

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

func init() {
	CreateCmd.Flags().StringVarP(&Name, "name", "n", "", "Task name")
	CreateCmd.Flags().StringVarP(&Description, "description", "d", "", "Task description")
	CreateCmd.Flags().StringSliceVarP(&Labels, "labels", "l", []string{}, "Task labels")
	CreateCmd.Flags().StringVarP(&Status, "status", "s", "to-do", "Task status")

	TaskCmd.AddCommand(CreateCmd, ListCmd, DeleteCmd, FinishCmd)
}

func Execute() {
	if err := TaskCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
